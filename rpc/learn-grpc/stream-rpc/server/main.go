package main

import (
	"google.golang.org/grpc"
	"learn-grpc/stream-rpc/pb"
	"log"
	"net"
	"strconv"
	"time"
)

type StreamServer struct {


}

func (s StreamServer) ListValue(request *pb.SimpleRequest, server pb.StreamServer_ListValueServer) error {
	for n := 0; n < 15; n++ {
		err := server.Send(&pb.StreamResponse{
			StreamValue:          request.Data + strconv.Itoa(n),

		})
		log.Println("send value:", request.Data + strconv.Itoa(n) )
		if err != nil {
			return err
		}
		time.Sleep(1 * time.Second)
	}
	return nil
}

const (
	Address string = ":8090"
)

func main() {
	listen, err := net.Listen("tcp", Address)
	if err != nil {
		log.Fatal("listen err:", err)
	}

	server := grpc.NewServer()
	pb.RegisterStreamServerServer(server, &StreamServer{})
	err = server.Serve(listen)
	if err != nil {
		log.Fatal("server run err:", err)
	}

}
