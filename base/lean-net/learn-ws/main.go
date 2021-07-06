package main

import (
	"context"
	"github.com/gorilla/websocket"
	"learn-go/base/lean-net/learn-ws/wsx"
	"log"
	"net/http"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}





func main() {
	http.HandleFunc("/ws", func(writer http.ResponseWriter, request *http.Request) {
		wsConn, err := upgrader.Upgrade(writer, request, nil)
		if err != nil {
			return
		}
		conn := wsx.NewWsConn(context.TODO(), wsConn, func(wsSocket wsx.Writer, data []byte) error {
			err := wsSocket.WriteJSON("123")
			return err
		}, func(err error, wsSocket wsx.Writer, cancelFunc context.CancelFunc) {
			log.Println("err = ", err)
			cancelFunc()
		})
		conn.Start()
	})
	err := http.ListenAndServe(":9001", http.DefaultServeMux)
	if err != nil {
		log.Fatalln(err)
	}
}

