package calculator

import (
	"context"

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
