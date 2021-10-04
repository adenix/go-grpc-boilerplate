package calculator

import (
	"context"
	"log"

	"github.com/adenix/go-grpc-boilerplate/gen/go/calculatorpb/v1"
)

type calculatorServiceServer struct {
	calculatorpb.UnimplementedCalculatorServiceServer
}

func NewCalculatorServiceServer() *calculatorServiceServer {
	return &calculatorServiceServer{}
}

func (s *calculatorServiceServer) Sum(ctx context.Context, in *calculatorpb.SumRequest) (*calculatorpb.SumResponse, error) {
	return &calculatorpb.SumResponse{
		Result: in.GetA() + in.GetB(),
	}, nil
}

func (g *calculatorServiceServer) Prime(in *calculatorpb.PrimeRequest, stream calculatorpb.CalculatorService_PrimeServer) error {
	var k uint64 = 2
	n := in.GetNumber()

	log.Printf("Executing prime number decomposition of %d", n)

	for n > 1 {
		if n%k == 0 {
			log.Printf("Sending divisor %d", k)
			stream.Send(&calculatorpb.PrimeResponse{Result: k})
			n = n / k
		} else {
			k++
			log.Printf("Divisor has increased to %d", k)
		}
	}

	return nil
}
