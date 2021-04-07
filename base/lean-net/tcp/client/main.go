package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os"
	"strings"
)

func main() {
	conn, err := net.Dial("tcp", "localhost:9090")
	if err != nil {
		log.Fatal("conn failed err:", err)
	}
	defer conn.Close()
	reader := bufio.NewReader(os.Stdin)
	for {
		input,err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("read from console failed,err:", err)
			break
		}
		trimInput := strings.TrimSpace(input)
		if input == "q" {
			break
		}
		_, err = conn.Write([]byte(trimInput))
		if err != nil {
			fmt.Println("write conn failed, err:", err)
			break
		}
	}
}

