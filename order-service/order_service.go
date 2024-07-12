package main

import (
	"context"
	"fmt"
	"net"

	//	pb "common/order"
	pb "generated-proto/proto/order"

	"google.golang.org/grpc"
)

type orderServiceServer struct {
	pb.UnimplementedOrderServiceServer
}

func (s *orderServiceServer) GetOrder(ctx context.Context, req *pb.OrderRequest) (*pb.OrderResponse, error) {
	order := req.OrderId
	return &pb.OrderResponse{OrderId: order}, nil
}

func main() {
	fmt.Println("order service is running on port 50052")

	lis, err := net.Listen("tcp", ":50052")
	if err != nil {
		panic(err)
	}
	grpcServer := grpc.NewServer()
	pb.RegisterOrderServiceServer(grpcServer, &orderServiceServer{})
	if err := grpcServer.Serve(lis); err != nil {
		panic(err)
	}
}
