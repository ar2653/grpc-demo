package main

import (
	"log"
	"net"

	pb "github.com/ar2653/grpc-sample-demo/proto"
	"google.golang.org/grpc"
)

const (
	port = ":50051"
)

type helloServer struct {
	pb.GreetServiceServer
}

func main() {
	lis, err := net.Listen("tcp", port)

	if err != nil {
		log.Fatalf("Failed to start the server %v", err)
	}

	grpcServer := grpc.NewServer()

	pb.RegisterGreetServiceServer(grpcServer, &helloServer{})

	log.Printf("Server started at port %v", lis.Addr())

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("Failed to serve %s", err)
	}
}
