package main

import (
	. "distributed/redisx"
	"github.com/go-redis/redis"
)

func main() {
	Client.Del("books")

	Client.ZAdd("books", redis.Z{
		Score:  9,
		Member: "java",
	}, redis.Z{
		Score:  8,
		Member: "go",
	}, redis.Z{
		Score:  7.2,
		Member: "python",
	}, redis.Z{
		Score:  5.5,
		Member: "z",
	})

	Out(Client.ZRange("books", 0, -1).Result()) // 按score顺序排

	Out(Client.ZRevRange("books", 0, -1).Result()) // 按score顺序排

	Out(Client.ZCard("books").Result()) // count

	Out(Client.ZScore("books", "java").Result()) // 4

	Out(Client.ZRank("books", "z").Result()) // 5 排名,从低到高

	Out(Client.ZRangeByScore("books", redis.ZRangeBy{
		Min:    "7",
		Max:    "8",
		Offset: 0,
		Count:  0,
	}).Result()) // 6

	Out(Client.ZRangeByScore("books", redis.ZRangeBy{
		Min:    "-inf", // inf 表示无穷大
		Max:    "8",
		Offset: 0,
		Count:  0,
	}).Result()) // 7

	Out(Client.ZRem("books", "java").Result()) // 8

}
