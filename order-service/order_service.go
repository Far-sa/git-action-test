package main

import (
	"context"
	"fmt"
	"net"

	//	pb "common/order"
	pb "common/genproto/common/protos/order"

	"google.golang.org/grpc"
)

type orderServiceServer struct {
	pb.UnimplementedOrderServiceServer
}

func (s *orderServiceServer) GetOrder(ctx context.Context, req *pb.GetOrderRequest) (*pb.GetOrderResponse, error) {
	return &pb.GetOrderResponse{Id: req.Id, Item: "book1"}, nil
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
