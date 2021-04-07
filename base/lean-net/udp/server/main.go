package main

import (
	"fmt"
	"log"
	"net"
)

func main() {
	listen, err := net.ListenUDP("udp", &net.UDPAddr{
		IP:   net.IPv4(0, 0, 0, 0),
		Port: 9090,
	})
	if err != nil {
		log.Fatal("listen failed, err=", err)
		return
	}

	for {
		var data [1024]byte
		n, addr, err := listen.ReadFromUDP(data[:])
		if err != nil {
			fmt.Println("read failed from addr, err=", err)
			break
		}

		go func() {
			fmt.Printf("addr: %v data: %v count: %v\n", addr,
			string(data[:]), n)
			_, err = listen.WriteToUDP([]byte("received success!"), addr)
			if err != nil {
				fmt.Printf("write failed ,err: %v\n", err)
			}
		}()
	}
}
