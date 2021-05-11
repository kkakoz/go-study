package main

import (
	"context"
	"google.golang.org/grpc"
	"io"
	"learn-grpc/client-stream-rpc/pb"
	"log"
	"strconv"
)

const (
	Address string = ":9001"
)

func main() {
	conn, err := grpc.Dial(Address, grpc.WithInsecure())
	if err != nil {
		log.Fatal("dial err:", err)
	}
	defer conn.Close()

	grpcClient := pb.NewClientStreamServerClient(conn)
	stream, err := grpcClient.RouteList(context.TODO())
	if err != nil {
		log.Fatal("call method err:", err)
	}

	for i := 0; i < 5; i++ {
		err := stream.Send(&pb.StreamRequest{
			StreamData: "hello" + strconv.Itoa(i),
		})
		//发送也要检测EOF，当服务端在消息没接收完前主动调用SendAndClose()关闭stream，此时客户端还执行Send()，则会返回EOF错误，所以这里需要加上io.EOF判断
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal("send err:", err)
		}
	}
	res, err := stream.CloseAndRecv()
	if err != nil {
		log.Fatal("recv res err:", err)
	}
	log.Println(res)
}
