package test

import (
	"encoding/json"
	"fmt"
	jsoniter "github.com/json-iterator/go"
	"testing"
)

type Task struct {
	Num int
}

type Job struct {
	Task
}

var MaxWorker = 5

type Worker struct {
	id int
	JobChannel chan Job
	WorkerPool chan chan Job
	exit chan bool
}

func TestABC(t *testing.T) {

	//pool := make(chan chan Job, MaxWorker)
	//
	//w := Worker{
	//	id: 1, JobChannel: make(chan Job), WorkerPool: pool, exit: make(chan bool),
	//}
	//
	//go func() {
	//	for {
	//		w.WorkerPool <- w.JobChannel
	//		fmt.Println("register private JobChannel to public WorkerPool", <- w.JobChannel)
	//	}
	//}()
	//
	//time.Sleep(10*time.Second)

}

func TestT(t *testing.T) {
	a := "123"
	marshal, _ := json.Marshal(a)
	fmt.Println(string(marshal))


}

type c func(a int)

func TestAA(t *testing.T)  {
	m := map[string]interface{}{
		"a": 1,
		"b":1.2,
		"c": "abc",
	}
	bytes, err := jsoniter.Marshal(&m)
	if err != nil {
		t.Fatal(err)
	}
	res := map[string]interface{}{}
	err = jsoniter.Unmarshal(bytes, &res)
	if err != nil {
		t.Fatal(err)
	}

}