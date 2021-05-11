package main

import (
	"context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"learn-grpc/grpc-deadlines/pb"
	"log"
	"net"
	"runtime"
	"time"
)

type StreamServer struct {


}

func (s StreamServer) Route(ctx context.Context, request *pb.SimpleRequest) (*pb.SimpleResponse, error) {
	data := make(chan *pb.SimpleResponse, 1)
	go handle(ctx, request, data)
	select {
	case res := <- data:
		return res, nil
	case <-ctx.Done():
		return nil, status.Errorf(codes.Canceled, "client called, abandoning")
	}
}

func handle(ctx context.Context, request *pb.SimpleRequest, data chan *pb.SimpleResponse) {
	select {
	case <-ctx.Done():
		log.Println(ctx.Err())
		runtime.Goexit() //超时后退出该Go协程
	case <-time.After(4 * time.Second): // 模拟耗时操作
		res := pb.SimpleResponse{
			Value: "hello " + request.Data,
		}
		// //修改数据库前进行超时判断
		// if ctx.Err() == context.Canceled{
		// 	...
		// 	//如果已经超时，则退出
		// }
		data <- &res
	}
}

const (
	Address string = ":8090"
)

func main() {
	listen, err := net.Listen("tcp", Address)
	if err != nil {
		log.Fatal("listen err:", err)
	}

	server := grpc.NewServer()
	pb.RegisterDeadlineServerServer(server, &StreamServer{})
	err = server.Serve(listen)
	if err != nil {
		log.Fatal("server run err:", err)
	}

}
