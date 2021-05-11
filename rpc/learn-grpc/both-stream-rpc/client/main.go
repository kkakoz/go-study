package main

import (
	"context"
	"google.golang.org/grpc"
	"io"
	"learn-grpc/both-stream-rpc/pb"
	"log"
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

	grpcClient := pb.NewBothStreamServerClient(conn)
	stream, err := grpcClient.Conversations(context.TODO())
	if err != nil {
		log.Fatal("call method err:", err)
	}

	for i := 0; i < 5; i++ {
		err := stream.Send(&pb.StreamRequest{Request: "hello"})
		if err != nil {
			log.Fatal("send err:", err)
		}
		res, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal("get stream err:", err)
		}
		log.Println("res = ", res)
	}
	err = stream.CloseSend()
	if err != nil {
		log.Fatal("recv res err:", err)
	}
}
