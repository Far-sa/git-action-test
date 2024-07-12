package main

import (
	"context"
	"log"
	"net/http"

	orderpb "common/genproto/common/protos/order"

	userpb "common/genproto/common/protos/user"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
)

func main() {
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	mux := runtime.NewServeMux()
	opts := []grpc.DialOption{grpc.WithInsecure()}

	err := userpb.RegisterUserServiceHandlerFromEndpoint(ctx, mux, "localhost:50051", opts)
	if err != nil {
		log.Fatalf("Failed to start HTTP gateway for UserService: %v", err)
	}

	err = orderpb.RegisterOrderServiceHandlerFromEndpoint(ctx, mux, "localhost:50052", opts)
	if err != nil {
		log.Fatalf("Failed to start HTTP gateway for OrderService: %v", err)
	}

	log.Println("Serving gRPC-Gateway on http://localhost:5000")
	if err := http.ListenAndServe(":5000", mux); err != nil {
		log.Fatalf("Failed to serve HTTP gateway: %v", err)
	}
}
