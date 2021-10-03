package gateway

import (
	"context"
	"fmt"
	"net/http"
	"strings"

	"github.com/adenix/go-grpc-boilerplate/gen/go/greetpb/v1"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
)

func Run(addr string) error {
	conn, err := grpc.DialContext(
		context.Background(),
		addr,
		grpc.WithInsecure(),
		grpc.WithBlock(),
	)
	if err != nil {
		return fmt.Errorf("failed to dial gRPC service: %w", err)
	}

	mux := runtime.NewServeMux()
	err = greetpb.RegisterGreetServiceHandler(context.Background(), mux, conn)
	if err != nil {
		return fmt.Errorf("failed to register REST gateway: %w", err)
	}

	srv := &http.Server{
		Addr: ":8080",
		Handler: http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			if strings.HasPrefix(r.URL.Path, "/api") {
				mux.ServeHTTP(w, r)
				return
			}
		}),
	}

	return srv.ListenAndServe()
}
