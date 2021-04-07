package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	rs1 := &RealServer{Addr:"127.0.0.1:2003"}
	rs1.Run()
	rs2 := &RealServer{Addr:"127.0.0.1:2004"}
	rs2.Run()

	quit := make(chan os.Signal)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
}

type RealServer struct {
	Addr string
}

func (r *RealServer) Run() {
	log.Println("starting httpserver at " + r.Addr)
	mux := http.NewServeMux()
	mux.HandleFunc("/", r.HelloHandler)
	mux.HandleFunc("/base/err", r.ErrHandler)
	server := &http.Server{
		Addr:    r.Addr,
		Handler: mux,
	}
	go func() {
		log.Fatal(server.ListenAndServe())
	}()
}

func (r *RealServer) HelloHandler(writer http.ResponseWriter, request *http.Request) {
	data := fmt.Sprintf("https://%s %s", r.Addr, request.URL.Path)
	realIp := fmt.Sprintf("remoteAddr=%s, X-Forwarded-For=%v,X-Real-Ip=%v\n",
		request.RemoteAddr, request.Header.Get("X-Forwarded-For"),
		request.Header.Get("X-Real-Ip"))
	io.WriteString(writer, data)
	io.WriteString(writer, realIp)
}

func (r *RealServer) ErrHandler(writer http.ResponseWriter, request *http.Request) {
	data := fmt.Sprintf("https://%s %s", r.Addr, request.URL.Path)
	realIp := fmt.Sprintf("remoteAddr=%s, X-Forwarded-For=%v,X-Real-Ip=%v\n",
		request.RemoteAddr, request.Header.Get("X-Forwarded-For"),
		request.Header.Get("X-Real-Ip"))
	io.WriteString(writer, data)
	io.WriteString(writer, realIp)
}
