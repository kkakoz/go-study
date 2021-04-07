package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
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
	proxy.ModifyResponse = func(res *http.Response) error {

		oldPayload, err := ioutil.ReadAll(res.Body)
		if err != nil {
			return err
		}
		newPayload := []byte("hello " + string(oldPayload))
		res.Body = ioutil.NopCloser(bytes.NewBuffer(newPayload))
		res.ContentLength = int64(len(newPayload))
		res.Header.Set("Content-Length", fmt.Sprint(len(newPayload)))
		return nil
	}
	log.Println("starting httpserver at " + addr)
	log.Fatal(http.ListenAndServe(addr, proxy))
}
