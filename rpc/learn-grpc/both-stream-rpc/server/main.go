package main

import (
	"google.golang.org/grpc"
	"io"
	"learn-grpc/both-stream-rpc/pb"
	"log"
	"net"
)

type StreamServer struct {
}

func (s StreamServer) Conversations(server pb.BothStreamServer_ConversationsServer) error {
	for {
		recv, err := server.Recv()
		if err == io.EOF {
			return nil
		}
		if err != nil {
			return err
		}
		err = server.Send(&pb.StreamResponse{
			Answer: "recv " + recv.Request,
		})
		if err != nil {
			return err
		}
		log.Println("recv data ", recv.Request)
	}
}

const (
	Address string = ":9001"
)

func main() {
	listen, err := net.Listen("tcp", Address)
	if err != nil {
		log.Fatal("listen err:", err)
	}

	server := grpc.NewServer()
	pb.RegisterBothStreamServerServer(server, &StreamServer{})
	err = server.Serve(listen)
	if err != nil {
		log.Fatal("server run err:", err)
	}

}
