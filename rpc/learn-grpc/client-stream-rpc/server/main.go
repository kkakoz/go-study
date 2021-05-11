package main

import (
	"google.golang.org/grpc"
	"io"
	"learn-grpc/client-stream-rpc/pb"
	"log"
	"net"
)

type StreamServer struct {
}

func (s StreamServer) RouteList(server pb.ClientStreamServer_RouteListServer) error {
	res := ""
	for {
		req, err := server.Recv()
		if err == io.EOF {
			return server.SendAndClose(&pb.SimpleResponse{
				Code:  1,
				Value: res,
			})
		}
		res += req.StreamData + "\n"
		if err != nil {
			return err
		}
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
	pb.RegisterClientStreamServerServer(server, &StreamServer{})
	err = server.Serve(listen)
	if err != nil {
		log.Fatal("server run err:", err)
	}

}
