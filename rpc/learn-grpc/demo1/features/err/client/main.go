package main

import (
	"context"
	"flag"
	"fmt"
	"google.golang.org/grpc"
	"learn-grpc/demo1/features/echopb/pb"
	"log"
	"os"
)

func main() {
	addr := flag.String("addr", "localhost:50051", "the address to connect to")
	flag.Parse()

	// Set up a connection to the server.
	conn, err := grpc.Dial(*addr, grpc.WithInsecure()) // To call service methods, we first need to create a gRPC channel to communicate with the server. We create this by passing the server address and port number to grpc.Dial()
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()

	c := pb.NewEchoClient(conn) // Once the gRPC channel is setup, we need a client stub to perform RPCs. We get this using the NewEchoClient method provided in the pb package we generated from our .proto.

	// Contact the server and print out its response.
	msg := "Madman hello world"
	if len(os.Args) > 1 {
		msg = os.Args[1]
	}

	resp, err := c.UnaryEcho(context.Background(), &pb.EchoRequest{Message: msg}) // Now let’s look at how we call our service methods. Note that in gRPC-Go, RPCs operate in a blocking/synchronous mode, which means that the RPC call waits for the server to respond, and will either return a response or an error.
	if err != nil {
		log.Fatalf("failed to call UnaryEcho: %v", err)
	}
	fmt.Printf("response:\n")
	fmt.Printf(" - %v\n", resp.GetMessage())
}
