package gateway

import (
	"context"
	"fmt"
	"io/fs"
	"log"
	"mime"
	"net/http"
	"strings"

	"github.com/adenix/go-grpc-boilerplate/gen"
	"github.com/adenix/go-grpc-boilerplate/gen/go/greetpb/v1"
	"github.com/adenix/go-grpc-boilerplate/third_party"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
)

func getSwaggerUIHandler() http.Handler {
	mime.AddExtensionType(".svg", "image/svg+xml")
	subFS, err := fs.Sub(third_party.SwaggerUI, "swagger_ui")
	if err != nil {
		log.Fatalf("failed to create sub filesystem for Swagger UI: %w", err)
	}
	return http.StripPrefix("/swagger-ui", http.FileServer(http.FS(subFS)))
}

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

	notFoundHandler := http.NotFoundHandler()

	swagger := getSwaggerUIHandler()
	srv := &http.Server{
		Addr: ":8080",
		Handler: http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			if strings.HasPrefix(r.URL.Path, "/api") {
				mux.ServeHTTP(w, r)
				return
			} else if r.URL.Path == "/v2/api-docs" {
				w.Header().Set("Content-Type", "application/json")
				w.Write(gen.OpenAPIV2)
				return
			} else if strings.HasPrefix(r.URL.Path, "/swagger-ui/") {
				swagger.ServeHTTP(w, r)
				return
			}
			notFoundHandler.ServeHTTP(w, r)
		}),
	}

	return srv.ListenAndServe()
}
