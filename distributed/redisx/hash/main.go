package main

import (
	. "distributed/redisx"
)

func main() {
	Client.Del("books")
	Client.HSet("books", "java", "think in java")

	Out(Client.HGet("books", "java").Result()) // 1

	Client.HSet("books", "java", "think java 2")

	Out(Client.HGet("books", "java").Result()) // 2

	Client.HSet("books", "go", "concurrency in go")

	Out(Client.HGetAll("books").Result())

	Out(Client.HLen("books"))

}
