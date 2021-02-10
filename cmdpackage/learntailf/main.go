package main

import (
	"fmt"
	"github.com/hpcloud/tail"
	"time"
)

func main() {
	fileName := "./my.log"
	config := tail.Config{
		ReOpen:    true,                                 //
		Follow:    true,                                 // 跟随重新打开文件
		Location:  &tail.SeekInfo{Offset: 0, Whence: 2}, //2表示从末尾读
		MustExist: false,                                //允许不存在
		Poll:      true, // 轮询方式
	}
	tails, err := tail.TailFile(fileName, config)
	if err != nil {
		panic(err)
	}
	for  {
		msg, ok := <-tails.Lines
		if !ok {
			fmt.Println("tail file close reopen, filename:", tails.Filename)
			time.Sleep(1 * time.Second)
			continue
		}
		fmt.Println("msg = ", msg.Text)
	}
}