package main

import (
	"context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"learn-grpc/grpc-deadlines/pb"
	"learn-grpc/grpc-security/pkg/auth"
	"log"
	"net"
)

type StreamServer struct {

}

func (s StreamServer) Route(ctx context.Context, request *pb.SimpleRequest) (*pb.SimpleResponse, error) {
	if err := auth.Check(ctx); err != nil {
		return nil, err
	}
	return &pb.SimpleResponse{Value: "hello" + request.Data}, nil
}

const (
	Address string = ":8090"
)

func main() {
	listen, err := net.Listen("tcp", Address)
	if err != nil {
		log.Fatal("listen err:", err)
	}
	creds, err := credentials.NewServerTLSFromFile("../pkg/tls/server.pem", "../pkg/tls/server.key")
	if err != nil {
		log.Fatalf("Failed to generate credentials %v", err)
	}
	var interceptor grpc.UnaryServerInterceptor
	interceptor = func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp interface{}, err error) {
		//拦截普通方法请求，验证Token
		err = auth.Check(ctx)
		if err != nil {
			return
		}
		// 继续处理请求
		return handler(ctx, req)
	}


	server := grpc.NewServer(grpc.Creds(creds), grpc.UnaryInterceptor(interceptor))
	pb.RegisterDeadlineServerServer(server, &StreamServer{})
	err = server.Serve(listen)
	if err != nil {
		log.Fatal("server run err:", err)
	}

}
