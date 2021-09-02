package wsx

import (
	"context"
	"github.com/gorilla/websocket"
)

type WsConn struct {
	wsSocket    *websocket.Conn
	readHandler ReadHandler
	errHandler  ErrHandler
	cancelFunc  context.CancelFunc
	ctx         context.Context
}

func NewWsConn(ctx context.Context, wsSocket *websocket.Conn, readHandler ReadHandler, errHandler ErrHandler) *WsConn {
	ctx, cancelFunc := context.WithCancel(ctx)
	return &WsConn{ctx: ctx, wsSocket: wsSocket, readHandler: readHandler, errHandler: errHandler, cancelFunc: cancelFunc}
}

type Writer interface {
	WriteJSON(v interface{}) error
	WriteMessage(messageType int, data []byte) error
}

func (w *WsConn) Start() {
	for {
		select {
		case <-w.ctx.Done():
			w.wsSocket.Close()
			return
		default:
			_, data, err := w.wsSocket.ReadMessage()
			if err != nil {
				w.errHandler(err, w.wsSocket, w.cancelFunc)
				continue
			}
			err = w.readHandler(w.wsSocket, data)
			if err != nil {
				w.errHandler(err, w.wsSocket, w.cancelFunc)
				continue
			}
		}
	}
}

type ReadHandler func(wsSocket Writer, data []byte) error

type ErrHandler func(err error, wsSocket Writer, cancelFunc context.CancelFunc)
