package main

import (
	"log"
	"net"

	"github.com/adenix/go-grpc-boilerplate/gateway"
	"github.com/adenix/go-grpc-boilerplate/gen/go/greetpb/v1"
	"github.com/adenix/go-grpc-boilerplate/greet"
	"google.golang.org/grpc"
)

func main() {
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("Failed to create listener on port 50051: %q ", err)
	}

	s := grpc.NewServer()
	greetpb.RegisterGreetServiceServer(s, greet.NewGreetServiceServer())

	log.Print("Starting gRPC service on :50051")
	go func() {
		if err = s.Serve(lis); err != nil {
			log.Fatalf("Failed to start gRPC service: %q", err)
		}
	}()

	log.Print("Starting REST service on :8080")
	if err = gateway.Run("dns:///0.0.0.0:50051"); err != nil {
		log.Fatalf("Failed to start REST service: %q", err)
	}
}
