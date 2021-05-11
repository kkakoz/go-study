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

	go func() {
		for {
			kv.Put(ctx, "test", "zhangsna")
			time.Sleep(1 * time.Second)
			kv.Delete(ctx, "test")
			time.Sleep(1 * time.Second)
		}
	}()
	watcher := clientv3.NewWatcher(client)
	watch := watcher.Watch(ctx, "test")
	go func() {
		for {
			select {
			case change := <- watch:
				for _, v := range change.Events {
					fmt.Println("type = ", v.Type)
					fmt.Println("now = ", string(v.Kv.Value))
					fmt.Println("pre = ", v.PrevKv)
				}
			}
		}
	}()
	select {

	}
}
