package main

import (
	"learn-go/base/lean-net/unpack"
	"log"
	"net"
)

func main() {
	conn, err := net.Dial("tcp", "localhost:9090")
	if err != nil {
		log.Fatal("conn failed err:", err)
	}
	defer conn.Close()
	err = unpack.Encode(conn, "hello world0")
	if err != nil {
		log.Fatal("write content err:", err)
	}
	err = unpack.Encode(conn, "hello world0")
	if err != nil {
		log.Fatal("write content err:", err)
	}
}
