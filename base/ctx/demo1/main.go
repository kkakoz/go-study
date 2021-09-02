package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	ctx, cancelFunc := context.WithCancel(context.TODO())
	funcs := make(chan context.CancelFunc, 1)
	funcs <- cancelFunc

	go func() {
		<-ctx.Done()
		fmt.Println("done")
	}()

	go func() {
		f := <- funcs
		fmt.Println("before cancel")
		f()
	}()
	time.Sleep(1 * time.Second)
}
