package main

import "distributed/redisx"

func main() {
	pipeline := redisx.Client.Pipeline()
	for i := 0; i < 10000; i++ {
		pipeline.PFAdd("users", i)
	}
	pipeline.Exec()

	redisx.Out(redisx.Client.PFCount("users").Result())

}
