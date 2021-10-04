package greet

import (
	"context"
	"fmt"
	"io"
	"log"
	"strings"
	"time"

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
	log.Print("Sending greeting")
	return &greetpb.GreetResponse{
		Result: fmt.Sprintf("Hello, %s!", formatName(in.GetGreeting())),
	}, nil
}

func (g *greetServiceServer) GreetManyTimes(in *greetpb.GreetManyTimesRequest, stream greetpb.GreetService_GreetManyTimesServer) error {
	log.Printf("GreetManyTimes procedure executed with %+v", in)
	for i := 0; i < 10; i++ {
		time.Sleep(1 * time.Second)
		log.Printf("Sending greeting #%d", i+1)
		stream.Send(&greetpb.GreetManyTimesResponse{
			Result: fmt.Sprintf("Hello, %s!", formatName(in.GetGreeting())),
			Count:  uint32(i + 1),
		})
	}
	return nil
}

func (g *greetServiceServer) LongGreet(stream greetpb.GreetService_LongGreetServer) error {
	log.Print("LongGreet procedure executed with a stream")

	names := ""
	for {
		req, err := stream.Recv()
		if err == io.EOF {
			log.Print("Sending greeting")
			return stream.SendAndClose(&greetpb.LongGreetResponse{
				Result: fmt.Sprintf("Hello, %s!", names),
			})
		} else if err != nil {
			log.Fatalf("Failed to read client stream: %w", err)
			return err
		}
		log.Printf("Recieved from stream: %+v", req)
		if names != "" {
			names += ", "
		}
		names += formatName(req.GetGreeting())
	}
}

func formatName(in *greetpb.Greeting) string {
	name := "World"
	if f := in.GetFirstName(); f != "" {
		name = strings.TrimSpace(f)

		if l := in.GetLastName(); l != "" {
			name += " " + strings.TrimSpace(l)
		}
	}
	return name
}
