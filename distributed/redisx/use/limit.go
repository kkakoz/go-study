package main

import (
	"distributed/redisx"
	"fmt"
	"github.com/go-redis/redis"
	"time"
)

func limit() {
	for i := 0; i < 15; i++ {
		fmt.Println(isActionAllowed("user1", "login", 60, 10))
		time.Sleep(50 * time.Millisecond)
	}
}

func isActionAllowed(id, action string, period, maxCount int64) bool {
	key := fmt.Sprintf("hist:%s:%s", id, action)
	t := time.Now().UnixMilli()
	pipe := redisx.Client.Pipeline()

	pipe.ZAdd(key, redis.Z{
		Score:  float64(t),
		Member: t,
	})

	pipe.ZRemRangeByScore(key, "0", fmt.Sprintf("%d", t-period*1000))

	pipe.ZCard(key)

	pipe.Expire(key, time.Duration(period+1)*time.Second)

	exec, err := pipe.Exec()
	if err != nil {
		return false
	}
	result, _ := exec[2].(*redis.IntCmd).Result()
	return result < maxCount
}
