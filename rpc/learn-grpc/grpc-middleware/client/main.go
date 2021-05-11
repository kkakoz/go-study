package main

import (
	"context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"learn-grpc/grpc-middleware/pb"
	"log"
	"time"
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

	grpcClient := pb.NewSimpleServerClient(conn)
	clientDeadline := time.Now().Add(time.Duration(3 * time.Second))
	ctx, _ := context.WithDeadline(context.TODO(), clientDeadline)
	res, err := grpcClient.Route(ctx, &pb.SimpleRequest{Data: "hello"})
	if err != nil {
		// 判断是否是超时错误
		status, ok := status.FromError(err)
		if ok {
			if status.Code() == codes.DeadlineExceeded {
				log.Fatal("time out err")
			}
		}
		log.Fatal("call method err:", err)
	}
	log.Println(res)
}
