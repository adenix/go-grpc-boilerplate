package main

import (
	"context"
	"io"
	"log"
	"strconv"
	"sync"
	"time"

	"github.com/adenix/go-grpc-boilerplate/gen/go/calculatorpb/v1"
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
	// calculator := calculatorpb.NewCalculatorServiceClient(conn)

	doUnary(c)
	doStream(c)
	doClientStreaming(c)
	// doPrime(calculator, 120)
	// doPrime(calculator, 12390392840)
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

func doStream(c greetpb.GreetServiceClient) {
	log.Print("Executing stream Greet procedure")

	req := &greetpb.GreetManyTimesRequest{
		Greeting: &greetpb.Greeting{
			FirstName: "Austin",
			LastName:  "Nicholas",
		},
	}

	resp, err := c.GreetManyTimes(context.Background(), req)
	if err != nil {
		log.Fatalf("Failed to execute GreetManyTimes in the gRPC service: %q", err)
	}

	var wg sync.WaitGroup
	for {
		msg, err := resp.Recv()
		if err == io.EOF {
			break
		}

		wg.Add(1)
		go func(msg *greetpb.GreetManyTimesResponse) {
			defer wg.Done()

			log.Printf("Recieved response #%d: %q", msg.GetCount(), msg.GetResult())
		}(msg)
	}

	wg.Wait()
}

func doClientStreaming(c greetpb.GreetServiceClient) {
	log.Print("Executing client stream Greet procedure")

	stream, err := c.LongGreet(context.Background())
	if err != nil {
		log.Fatalf("Failed to exedcute LongGreet in gRPC service: %q", err)
	}

	log.Print("Sending greeting request #1")
	err = stream.Send(&greetpb.LongGreetRequest{
		Greeting: &greetpb.Greeting{
			FirstName: "Austin",
			LastName:  "Nicholas",
		},
	})
	if err != nil {
		log.Fatalf("Failed to send message to LongGreet: %q", err)
	}

	time.Sleep(2 * time.Second)

	log.Print("Sending greeting request #2")
	err = stream.Send(&greetpb.LongGreetRequest{
		Greeting: &greetpb.Greeting{
			FirstName: "Adam",
			LastName:  "MacDonald-Corrao",
		},
	})
	if err != nil {
		log.Fatalf("Failed to send message to LongGreet: %q", err)
	}

	if resp, err := stream.CloseAndRecv(); err != nil {
		log.Fatalf("Failed to recieve response: %q", err)
	} else {
		log.Printf("Recieved response: %q", resp.GetResult())
	}
}

func doPrime(c calculatorpb.CalculatorServiceClient, n uint64) {
	log.Print("Executing prime number procedure")

	req := &calculatorpb.PrimeRequest{
		Number: n,
	}

	resp, err := c.Prime(context.Background(), req)
	if err != nil {
		log.Fatalf("Failed to execute Prime in the gRPC service: %q", err)
	}

	register := ""
	for {
		msg, err := resp.Recv()
		if err == io.EOF {
			break
		}

		if register != "" {
			register += "*"
		}
		register += strconv.FormatUint(msg.GetResult(), 10)
	}

	log.Printf("Recieved result: %d=%s", n, register)
}
