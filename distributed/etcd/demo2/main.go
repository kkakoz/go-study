package main

import (
	"context"
	"fmt"
	"github.com/coreos/etcd/clientv3"
	"time"
)

func main() {
	config := clientv3.Config{
		Endpoints:   []string{"47.98.156.22:2379"},
		DialTimeout: 5 * time.Second,
	}
	client, err := clientv3.New(config)
	if err != nil {
		fmt.Println(err)
		return
	}
	ctx, _ := context.WithCancel(context.TODO())
	kv := clientv3.NewKV(client)
	lease := clientv3.NewLease(client)

	grant, err := lease.Grant(ctx, 10)
	if err != nil {
		fmt.Println(err)
		return
	}
	leaseID := grant.ID
	_, err = kv.Put(ctx, "temp", "lisi", clientv3.WithLease(leaseID))
	if err != nil {
		fmt.Println(err)
		return
	}
	for {
		getres, err := kv.Get(ctx, "temp")
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Println(getres.Kvs)
		if getres.Count < 1 {
			fmt.Println("过期了")
			return
		}
		time.Sleep(2 * time.Second)
	}

}