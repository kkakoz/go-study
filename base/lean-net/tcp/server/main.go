package main

import (
	"fmt"
	"log"
	"net"
)

func main() {
	listen, err := net.Listen("tcp", "0.0.0.0:9090")
	if err != nil {
		log.Fatal("listen err:", err)
	}
	defer listen.Close()

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
		bytes := make([]byte, 1024)
		_, err := conn.Read(bytes)
		if err != nil {
			fmt.Printf("read from conn failed, err:", err)
			break
		}
		str := string(bytes)
		fmt.Printf("receive from client, data: %v\n", str)
	}
}
