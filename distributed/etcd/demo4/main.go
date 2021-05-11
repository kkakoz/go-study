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
	put := clientv3.OpPut("name", "ttt")
	res, err := kv.Do(ctx, put)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(res.Put().Header.Revision)
}