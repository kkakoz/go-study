package main

import (
	"log"
	"net/http"
	"time"
)

func main() {
	// 路由器
	mux := http.NewServeMux()
	// 设置路由
	mux.HandleFunc("/bye", sayBye)
	// 创建服务器
	server := http.Server{
		Addr: ":9090",
		WriteTimeout: time.Second * 3,
		Handler: mux,
	}
	log.Println("start http server")
	log.Fatal(server.ListenAndServe())
}

func sayBye(writer http.ResponseWriter, request *http.Request) {
	time.Sleep(1 * time.Second)
	writer.Write([]byte("bye"))
}
