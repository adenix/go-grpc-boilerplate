package main

import (
	"log"

	"github.com/adenix/go-grpc-boilerplate/gen/go/greetpb/v1"
	"google.golang.org/grpc"
)

func main() {
	conn, err := grpc.Dial(":50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Failed to connect to gRPC service on port 50051: %q", err)
	}
	defer conn.Close()

	log.Print("Creating Greet Service Client")
	c := greetpb.NewGreetServiceClient(conn)

	log.Printf("Client created: %f", c)
}
