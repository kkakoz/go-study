package main

import (
	. "distributed/redisx"
	"fmt"
	"time"
)

var count = 1

func out(key string) {
	result, _ := Client.Get(key).Result()
	fmt.Println(count, "=", result)
	count++
}

func main() {
	out("name") // 1

	Client.Set("name", "张三", 0)

	out("name") // 2

	Client.Set("name", "张三", 1*time.Second)

	out("name") // 3

	time.Sleep(1 * time.Second)

	out("name") // 4

	Client.Set("name", "张三", 1*time.Second)

	out("name") // 5

	Client.Del("name")

	out("name") // 6

	fmt.Println(Client.SetNX("key1", "v1", 1*time.Second).Result())

	out("key1") // 7

	fmt.Println(Client.SetNX("key1", "v2", 1*time.Second).Result())

	Client.Set("num", 1, 0)

	Client.Incr("num")

	out("num") // 8
}
