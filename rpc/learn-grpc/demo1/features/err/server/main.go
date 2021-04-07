package main

import (
	"context"
	"flag"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"learn-grpc/demo1/features/echopb/pb"
	"log"
	"net"
)

type server struct {}

func (s server) UnaryEcho(ctx context.Context, req *pb.EchoRequest) (*pb.EchoResponse, error) {
	if len(req.Message) > 10 {
		return nil, status.Errorf(codes.Unimplemented, "length of message cannot be more than 10")
	}
	return &pb.EchoResponse{Message: req.GetMessage()}, nil
}

func (s server) ServerStreamingEcho(*pb.EchoRequest, pb.Echo_ServerStreamingEchoServer) error {
	return  status.Errorf(codes.Unimplemented, "method ServerStreamingEcho not implemented")
}

func (s server) ClientStreamEcho(pb.Echo_ClientStreamEchoServer) error {
	return status.Errorf(codes.Unimplemented, "method ClientStreamEcho not implemented")
}

func (s server) BidirectionalStreamingEcho(pb.Echo_BidirectionalStreamingEchoServer) error {
	return status.Errorf(codes.Unimplemented, "method BidirectionalStreamingEcho not implemented")
}

func main() {
	port := flag.Int("port", 50051, "the port to serve on")
	flag.Parse()

	listen, err := net.Listen("tcp", fmt.Sprintf(":%d", *port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	fmt.Printf("server listening at %v\n", listen.Addr())

	s := grpc.NewServer()
	pb.RegisterEchoServer(s, &server{})
	if err := s.Serve(listen); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}

