package main

import (
	"context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"learn-grpc/grpc-deadlines/pb"
	"learn-grpc/grpc-security/pkg/auth"
	"log"
)

const (
	Address string = ":8090"
)

func main() {
	creds, err := credentials.NewClientTLSFromFile("../pkg/tls/server.pem", "go-grpc-example")
	conn, err := grpc.Dial(Address,
		grpc.WithTransportCredentials(creds),
		grpc.WithPerRPCCredentials(&auth.Token{
			AppID:     "grpc_token",
			AppSecret: "123456",
		}),
	)
	if err != nil {
		log.Fatal("dial err:", err)
	}
	defer conn.Close()

	grpcClient := pb.NewDeadlineServerClient(conn)
	res, err := grpcClient.Route(context.TODO(), &pb.SimpleRequest{Data: "hello"})
	if err != nil {
		log.Fatal("call method err:", err)
	}
	log.Println(res)
}
