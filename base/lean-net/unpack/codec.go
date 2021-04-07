/*
	解决粘包和拆包
 */
package unpack

import (
	"encoding/binary"
	"errors"
	"io"
)

const MsgHeader = "12345678"

func Encode(bytesBuffer io.Writer, content string) error {
	err := binary.Write(bytesBuffer, binary.BigEndian, []byte(MsgHeader))
	if err != nil {
		return err
	}
	clen := int32(len([]byte(content)))
	err = binary.Write(bytesBuffer, binary.BigEndian, clen)
	if err != nil {
		return err
	}
	binary.Write(bytesBuffer, binary.BigEndian, []byte(content))
	if err != nil {
		return err
	}
	return nil
}

func Decode(bytesBuffer io.Reader) ([]byte, error) {
	magicBuf := make([]byte, len(MsgHeader))
	_, err := io.ReadFull(bytesBuffer, magicBuf)
	if err != nil {
		return nil, err
	}
	if string(magicBuf) != MsgHeader {
		return nil, errors.New("msg_header error")
	}
	lengthBuf := make([]byte, 4)
	_, err = io.ReadFull(bytesBuffer, lengthBuf)
	if err != nil {
		return nil, err
	}
	length := binary.BigEndian.Uint32(lengthBuf)
	bodyBuf := make([]byte, length)
	_, err = io.ReadFull(bytesBuffer, bodyBuf)
	if err != nil {
		return nil, err
	}
	return bodyBuf, nil
}
