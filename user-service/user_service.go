package main

import (
	"context"
	"fmt"
	"net"

	pb "common/genproto/common/protos/user"

	"google.golang.org/grpc"
)

type userServiceServer struct {
	pb.UnimplementedUserServiceServer
}

func (s *userServiceServer) GetUser(ctx context.Context, req *pb.GetUserRequest) (*pb.GetUserResponse, error) {
	return &pb.GetUserResponse{Id: req.Id, Name: "John Doe"}, nil
}

func main() {
	fmt.Println("user service is running on port 50051")

	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		panic(err)
	}
	grpcServer := grpc.NewServer()
	pb.RegisterUserServiceServer(grpcServer, &userServiceServer{})
	if err := grpcServer.Serve(lis); err != nil {
		panic(err)
	}
}
