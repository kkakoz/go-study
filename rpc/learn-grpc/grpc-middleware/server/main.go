package main

import (
	"context"
	grpc_middleware "github.com/grpc-ecosystem/go-grpc-middleware"
	grpc_zap "github.com/grpc-ecosystem/go-grpc-middleware/logging/zap"
	"google.golang.org/grpc"
	"learn-grpc/grpc-middleware/pb"
	"learn-grpc/grpc-middleware/server/middleware"
	"log"
	"net"
)

type StreamServer struct {


}

func (s StreamServer) Route(ctx context.Context, request *pb.SimpleRequest) (*pb.SimpleResponse, error) {
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

	server := grpc.NewServer(
		grpc.StreamInterceptor(grpc_middleware.ChainStreamServer(
			grpc_zap.StreamServerInterceptor(middleware.ZapInterceptor()),
			)),
		grpc.UnaryInterceptor(grpc_middleware.ChainUnaryServer(
			grpc_zap.UnaryServerInterceptor(middleware.ZapInterceptor()),
		)))
	pb.RegisterSimpleServerServer(server, &StreamServer{})
	err = server.Serve(listen)
	if err != nil {
		log.Fatal("server run err:", err)
	}

}
