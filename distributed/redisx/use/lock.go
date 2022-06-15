package main

import (
	"distributed/redisx"
	"time"
)

func lockRun() {
	redisx.Out(redisx.Client.SetNX("lock1", true, time.Second).Result())

	redisx.Out(redisx.Client.SetNX("lock1", true, time.Second).Result())
}

func lock(key string) bool {
	result, err := redisx.Client.SetNX(key, true, time.Second).Result()
	if err != nil {
		return false
	}
	return result
}
