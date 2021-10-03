package greet

import (
	"context"
	"fmt"
	"log"
	"strings"

	"github.com/adenix/go-grpc-boilerplate/gen/go/greetpb/v1"
)

type greetServiceServer struct {
	greetpb.UnimplementedGreetServiceServer
}

func NewGreetServiceServer() *greetServiceServer {
	return &greetServiceServer{}
}

func (g *greetServiceServer) Greet(ctx context.Context, in *greetpb.GreetRequest) (*greetpb.GreetResponse, error) {
	log.Printf("Greet procedure executed with %+v", in)

	name := "World"
	if f := in.GetGreeting().GetFirstName(); f != "" {
		name = strings.TrimSpace(f)

		if l := in.GetGreeting().GetLastName(); l != "" {
			name += " " + strings.TrimSpace(l)
		}
	}

	return &greetpb.GreetResponse{
		Result: fmt.Sprintf("Hello, %s!", name),
	}, nil
}
