package main

import (
	"bufio"
	"log"
	"net/http"
	"net/url"
)

func main() {
	http.HandleFunc("/", handler)
	err := http.ListenAndServe(":" + port, nil)
	if err != nil {
		log.Fatal(err)
	}
}

var (
	proxyAddr = "http://127.0.0.1:2003"
	port      = "2002"
)

func handler(wr http.ResponseWriter, req *http.Request) {
	proxy, err := url.Parse(proxyAddr)
	if err != nil {
		wr.WriteHeader(http.StatusBadGateway)
		return
	}
	req.URL.Scheme = proxy.Scheme
	req.URL.Host = proxy.Host

	// 请求下游
	transport := http.DefaultTransport
	res, err := transport.RoundTrip(req)
	if err != nil {
		log.Println(err)
		return
	}

	for k, value := range res.Header {
		for _, v := range value {
			wr.Header().Add(k, v)
		}
	}
	defer res.Body.Close()
	_, err = bufio.NewReader(res.Body).WriteTo(wr)
	if err != nil {
		wr.WriteHeader(http.StatusBadGateway)
	}
}
