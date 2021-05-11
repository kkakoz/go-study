package main

import (
	"context"
	"google.golang.org/grpc"
	"io"
	"learn-grpc/stream-rpc/pb"
	"log"
)

const (
	Address string = ":8090"
)

func main() {
	conn, err := grpc.Dial(Address, grpc.WithInsecure())
	if err != nil {
		log.Fatal("dial err:", err)
	}
	defer conn.Close()

	grpcClient := pb.NewStreamServerClient(conn)
	stream, err := grpcClient.ListValue(context.TODO(), &pb.SimpleRequest{
		Data: "hello",
	})
	if err != nil {
		log.Fatal("call method err:", err)
	}
	for {
		// 默认每次接收最大长度为4M
		recv, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal("get stream err", recv)
		}
		log.Println(recv)
	}
	//可以使用CloseSend()关闭stream，这样服务端就不会继续产生流消息
	//调用CloseSend()后，若继续调用Recv()，会重新激活stream，接着之前结果获取消息
	stream.CloseSend()
}
