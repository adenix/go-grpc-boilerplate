package main

import (
	"log"
	"net"

	"github.com/adenix/go-grpc-boilerplate/gen/go/greetpb/v1"
	"github.com/adenix/go-grpc-boilerplate/server"
	"google.golang.org/grpc"
)

func main() {
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("Failed to create listener on port 50051: %q ", err)
	}

	s := grpc.NewServer()
	greetpb.RegisterGreetServiceServer(s, server.New())

	log.Print("Starting gRPC service on :50051")
	if err = s.Serve(lis); err != nil {
		log.Fatalf("Failed to start gRPC service: %q", err)
	}
}
