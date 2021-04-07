package main

import (
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
)

var addr = "127.0.0.1:2002"

func main() {
	// 请求 127.0.0.1/xxx => 127.0.0.1/base/xxx
	rs1 := "http://127.0.0.1:2003/base"
	url1, err := url.Parse(rs1)
	if err != nil {
		log.Println(err)
	}
	proxy := httputil.NewSingleHostReverseProxy(url1)
	log.Println("starting httpserver at " + addr)
	log.Fatal(http.ListenAndServe(addr, proxy))
}
