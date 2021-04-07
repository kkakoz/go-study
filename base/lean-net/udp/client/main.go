package main

import (
	"fmt"
	"log"
	"net"
)

func main() {
	conn, err := net.DialUDP("udp", nil, &net.UDPAddr{
		IP:   net.IPv4(0, 0, 0, 0),
		Port: 9090,
	})

	if err != nil {
		log.Fatal("connect failed, err: ", err)
	}

	for i := 0; i < 100; i++ {
		_, err := conn.Write([]byte("hello server"))
		if err != nil {
			fmt.Printf("send data failed err:", err)
			return
		}

		result := make([]byte, 1024)
		n, remoteAddr, err := conn.ReadFromUDP(result)
		if err != nil {
			fmt.Printf("received data failed, err:", err)
			return
		}
		fmt.Printf("received from addr: %v data: %v\n", remoteAddr, string(result[:n]))
	}
}
