package main

import (
	"fmt"
	"net/http"
	_ "net/http/pprof"
	"time"
)

func main() {
	go func() {
		for  {
			time.Sleep(1 * time.Second)
			fmt.Println("for loop")
		}
	}()
	http.ListenAndServe(":9090", nil)
}
