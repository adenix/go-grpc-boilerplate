package main

import (
	"context"
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

	doUnary(c)
}

func doUnary(c greetpb.GreetServiceClient) {
	log.Print("Executing unary Greet procedure")

	req := &greetpb.GreetRequest{
		Greeting: &greetpb.Greeting{
			FirstName: "Austin",
			LastName:  "Nicholas",
		},
	}
	resp, err := c.Greet(context.Background(), req)
	if err != nil {
		log.Fatalf("Failed to execute Greet in the gRPC service: %q", err)
	}

	log.Printf("Recieved response: %q", resp.GetResult())
}
