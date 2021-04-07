package main

import (
	"fmt"
	"io"
	"net"
	"net/http"
	"strings"
	"time"
)

type Pxy struct {

}

func (p *Pxy) ServeHTTP(rw http.ResponseWriter, req *http.Request) {
	fmt.Printf("received request %s %s %s\n",req.Method, req.Host,
		req.RemoteAddr)
	transport := http.DefaultTransport
	outReq := new(http.Request)
	*outReq = *req
	if clientIP, _, err := net.SplitHostPort(req.RemoteAddr); err!=nil {
		if prior, ok := outReq.Header["X-Forwarded-For"]; ok {
			clientIP = strings.Join(prior, ", ") + ", " + clientIP
		}
		outReq.Header.Set("X-Forwarded-For", clientIP)
	}
	// 请求下游
	res, err := transport.RoundTrip(outReq)
	if err != nil {
		rw.WriteHeader(http.StatusBadGateway)
		return
	}
	defer res.Body.Close()
	// 返回请求给上游
	for key, value := range res.Header {
		for _, v := range value {
			rw.Header().Add(key, v)
		}
	}
	rw.WriteHeader(res.StatusCode)
	_, err = io.Copy(rw, res.Body)
	if err != nil {
		rw.WriteHeader(http.StatusBadGateway)
		return
	}
	return
}

func main() {
	server := http.Server{
		Addr: ":9090",
		WriteTimeout: time.Second * 3,
		Handler: &Pxy{},
	}
	err := server.ListenAndServe()
	if err != nil {
		panic(err)
	}
}
