package server

import "github.com/adenix/go-grpc-boilerplate/gen/go/greetpb/v1"

type server struct {
	greetpb.UnimplementedGreetServiceServer
}

func New() *server {
	return &server{}
}
