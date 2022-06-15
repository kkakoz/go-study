package main

import (
	"distributed/redisx"
	"fmt"
	"github.com/go-redis/redis"
	"log"
	"time"
)

func lateQueueRun() {
	go lateQueue()
	redisx.Client.ZAdd("delay-queue", redis.Z{
		Score:  float64(time.Now().Unix() + 5),
		Member: "task1",
	})
	redisx.Client.ZAdd("delay-queue", redis.Z{
		Score:  float64(time.Now().Unix() + 10),
		Member: "task2",
	})
	select {}
}

func lateQueue() {
	for {
		result, err := redisx.Client.ZRangeByScore("delay-queue", redis.ZRangeBy{
			Min:    "0",
			Max:    fmt.Sprintf("%d", time.Now().Unix()),
			Offset: 0,
			Count:  1,
		}).Result()
		if err != nil {
			log.Fatal(err)
		}
		if len(result) > 0 {
			i, err := redisx.Client.ZRem("delay-queue", result[0]).Result()
			if err != nil {
				log.Fatal(err)
			}
			if i > 0 {
				fmt.Println("success = ", result[0])
			}
		} else {
			fmt.Println("not get ")
			time.Sleep(1 * time.Second)
		}

	}

}
