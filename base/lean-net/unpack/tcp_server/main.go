package main

import (
	"fmt"
	"learn-go/base/lean-net/unpack"
	"log"
	"net"
)

func main() {
	listen, err := net.Listen("tcp", "0.0.0.0:9090")
	if err != nil {
		log.Fatal("listen err:", err)
	}

	for {
		conn, err := listen.Accept()
		if err != nil {
			log.Println("accept fail, err:", err)
			continue
		}

		go process(conn)

	}
}

func process(conn net.Conn) {
	defer conn.Close()
	for {
		content, err := unpack.Decode(conn)
		if err != nil {
			fmt.Printf("read from conn failed, err:", err)
			break
		}
		str := string(content)
		fmt.Printf("receive from client, data: %v\n", str)
	}
}
