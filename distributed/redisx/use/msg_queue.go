package main

import (
	"distributed/redisx"
	"math/rand"
	"time"
)

func msgQueue(key string) {
	go push(key)
	go pop(key)
	select {}
}

func push(key string) {
	for {
		redisx.Client.LPush(key, rand.Int())
		time.Sleep(5 * time.Second)
	}

}

func pop(key string) {
	for {
		// 长时间没有数据可能断开连接
		redisx.Out(redisx.Client.BLPop(0, key).Result())
	}

}
