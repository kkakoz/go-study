package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net"
	"net/http"
	"time"
)

func main() {
	// 创建连接池

	transprot := &http.Transport{
		DialContext: (&net.Dialer{
			Timeout:   30 * time.Second, // 连接超时
			KeepAlive: 30 * time.Second, // 探活时间
		}).DialContext,
		MaxConnsPerHost:       100, // 最大空闲连接
		IdleConnTimeout:       90 * time.Second, // 空闲超过时间
		TLSHandshakeTimeout:   10 * time.Second, // tls握手超时时间
		ExpectContinueTimeout: 1 * time.Second, // continue状态码超时时间
	}

	client := &http.Client{
		Timeout:   time.Second * 3,
		Transport: transprot,
	}

	res, err := client.Get("http://127.0.0.1:9090/bye")
	if err != nil {
		log.Fatal("http client get err:", err)
	}
	defer res.Body.Close()
	bds, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Fatal("read body err:", err)
	}
	fmt.Println(string(bds))
}
