package main

import (
	. "distributed/redisx"
)

func main() {
	Client.Del("books")

	Out(Client.SAdd("books", "python").Result()) // 1

	Out(Client.SAdd("books", "python").Result()) // 2

	Out(Client.SAdd("books", "go").Result()) // 3

	Out(Client.SMembers("books")) // 4

	Out(Client.SIsMember("books", "go").Result()) // 5

	Out(Client.SCard("books").Result()) // 6 count

	Out(Client.SPop("books").Result()) // 7
}
