package main

import (
	pb "common/genproto/common/protos/order"
	"context"
	"fmt"
	"net"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"google.golang.org/grpc"
)

// Helper function to setup gRPC server and client
func setupMockServerAndClient(t *testing.T) (pb.OrderServiceClient, func()) {
	lis, err := net.Listen("tcp", ":0") // Listen on any available port
	require.NoError(t, err, "Failed to listen")

	grpcServer := grpc.NewServer()
	pb.RegisterOrderServiceServer(grpcServer, &orderServiceServer{})
	go grpcServer.Serve(lis)

	conn, err := grpc.Dial(fmt.Sprintf(":%d", lis.Addr().(*net.TCPAddr).Port), grpc.WithInsecure())
	require.NoError(t, err, "Failed to dial")

	client := pb.NewOrderServiceClient(conn)

	cleanup := func() {
		conn.Close()
		grpcServer.Stop()
		lis.Close()
	}

	return client, cleanup
}

func TestGetOrder(t *testing.T) {

	type testCase struct {
		name     string
		req      *pb.GetOrderRequest
		wantID   string
		wantItem string
		wantErr  bool
	}

	cases := []testCase{
		{"ValidRequest", &pb.GetOrderRequest{Id: "123"}, "123", "book1", false},
		{"InvalidRequest", &pb.GetOrderRequest{}, "", "", true},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			ctx := context.Background()
			server := &orderServiceServer{}
			resp, err := server.GetOrder(ctx, tc.req)

			if tc.wantErr {
				assert.Error(t, err, "Expected error")
			} else {
				assert.NoError(t, err, "Unexpected error")
				assert.Equal(t, tc.wantID, resp.Id, "Expected and actual IDs do not match")
				assert.Equal(t, tc.wantItem, resp.Item, "Expected and actual items do not match")
			}
		})
	}
}

func TestIntegrationGetOrder(t *testing.T) {
	client, cleanup := setupMockServerAndClient(t)
	defer cleanup()

	type testCase struct {
		name     string
		req      *pb.GetOrderRequest
		wantID   string
		wantItem string
		wantErr  bool
	}

	cases := []testCase{
		{"ValidRequest", &pb.GetOrderRequest{Id: "456"}, "456", "book1", false},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			resp, err := client.GetOrder(context.Background(), tc.req)

			if tc.wantErr {
				assert.Error(t, err, "Expected error")
			} else {
				assert.NoError(t, err, "Unexpected error")
				assert.Equal(t, tc.wantID, resp.Id, "Expected and actual IDs do not match")
				assert.Equal(t, tc.wantItem, resp.Item, "Expected and actual items do not match")
			}
		})
	}
}
