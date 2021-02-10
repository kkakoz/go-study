package main

import (
	"context"
	"github.com/micro/go-micro/v2/client"
	"log"
	pb "hello/proto/hello"
)

func main() {
	testCallFunc()
}

func testCallFunc(){
	// 获取hello服务
	// 这里第一个参数"go.micro.service.hello"必须与hello-service注册信息一致
	// 一般由micro生成的项目默认服务名为：{namespace 默认[go.micro]}.{type 默认[service]}.{项目名}组成
	// 如果要修改默认值，在生成项目时可以这样： micro --namespace=XXX --type=YYYY ZZZZ
	// 当然也可以直接修改main.go中micro.Name("go.micro.service.hello")的内容
	helloService := pb.NewHelloService("go.micro.service.hello", client.DefaultClient)

	// 默认生成的hello服务中自带三个接口: Call() Stream() PingPong(),分别对应参数调用、流传输和心跳
	resp, err := helloService.Call(context.Background(), &pb.Request{
		Name: "xiaoxie",
	})
	if err != nil {
		log.Panic("call func", err)
	}
	log.Println("call func success!", resp.Msg)
}
