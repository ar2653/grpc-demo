package main

import (
	"context"
	"io"
	"log"
	"time"

	pb "github.com/ar2653/grpc-sample-demo/proto"
)

func CallSayHelloBidirectionalStreaming(client pb.GreetServiceClient, names *pb.NamesList) {
	log.Printf("Bidirectional stream started")
	stream, err := client.SayHelloBidirectionalStreaming(context.Background())
	if err != nil {
		log.Fatalf("Could not send names: %v", err)
	}

	c := make(chan struct{})

	go func() {
		for {
			message, err := stream.Recv()
			if err == io.EOF {
				break
			}
			if err != nil {
				log.Fatalf("Error while streaming: %v", err)
			}
			log.Println(message)
		}
		close(c)
	}()

	for _, name := range names.Names {
		req := &pb.HelloRequest{
			Name: name,
		}
		if err := stream.Send(req); err != nil {
			log.Fatalf("Error while sending the request: %v", err)
		}
		time.Sleep(2 * time.Second)
	}
	stream.CloseSend()
	<-c
	log.Printf("Bidirectional streaming finished")
}
