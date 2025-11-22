package grpc

import (
	"log"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	pb "fitness-app-microservices/api-gateway/proto"
)

// NewUserClient creates a gRPC client for the UserService
func NewUserClient(addr string) pb.UserServiceClient {
	conn, err := grpc.NewClient(addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("Failed to connect to user service: %v", err)
	}
	return pb.NewUserServiceClient(conn)
}
