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
	ctx, cancel := context.WithCancel(context.TODO())

	lease := clientv3.NewLease(client)
	kv := clientv3.NewKV(client)
	grant, err := lease.Grant(ctx, 5)
	if err != nil {
		fmt.Println(err)
		return
	}

	leaseId := grant.ID
	watcher := clientv3.NewWatcher(client)
	watch := watcher.Watch(context.TODO(), "lock2")
	go func() {
		for v := range watch {
			fmt.Println("watch")
			fmt.Println(v.Events[0].Type)
			fmt.Println(v.Header.Revision)
			fmt.Println(v.CompactRevision)
		}
		fmt.Println("end watch")
	}()


	_, err = kv.Put(context.TODO(), "lock2", "true", clientv3.WithLease(leaseId))
	if err != nil {
		fmt.Println(err)
		return
	}
	keepChan, err := lease.KeepAlive(ctx, leaseId)
	if err != nil {
		fmt.Println(err)
		return
	}

	go func() {
		for keepRes := range keepChan{
			fmt.Println("自动续租应答", keepRes.ID)
		}
		_, err = lease.Revoke(context.TODO(), leaseId)
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Println("end keepres")
	}()


	time.Sleep(time.Second * 12)
	cancel()
	fmt.Println("取消续租")
	time.Sleep(time.Second * 5)

	//txn := kv.Txn(context.TODO())
	//// 如果key不存在
	//txn.If(clientv3.Compare(clientv3.CreateRevision("lock2"),"=", 0)).
	//	Then(clientv3.OpPut("lock2", "", clientv3.WithLease(leaseId))).
	//	Else(clientv3.OpGet("lock2"))
	//txnRes, err := txn.Commit()
	//if err != nil {
	//	fmt.Println(err)
	//	return
	//}
	//if !txnRes.Succeeded {
	//	fmt.Println("没有抢到锁:", txnRes.Responses[0].GetResponseRange())
	//	return
	//}
	//fmt.Println("抢到锁了")
}

