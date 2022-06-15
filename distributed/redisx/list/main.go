package main

import (
	. "distributed/redisx"
)

func main() {
	Client.Del("users")

	Client.LPush("users", "李四")

	Out(Client.LLen("users")) // 1

	Out(Client.LPop("users").Result()) // 2

	Out(Client.LPop("users").Result()) // 3

	Client.Del("nums")

	Client.LPush("nums", 1, 2, 3, 4, 5)

	Out(Client.LIndex("nums", 2).Result()) // 4

	Out(Client.LIndex("nums", -1).Result()) // 5

	Out(Client.LTrim("nums", 2, 4).Result()) // 6 保留 3 2 1 o(n)

	Out(Client.LIndex("nums", 0).Result()) // 7

	Out(Client.LRange("nums", 0, 2).Result()) // 8  o(n)

}
