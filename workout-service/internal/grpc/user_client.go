package grpc

import (
	"context"
	"log"

	pb "fitness-app-microservices/proto"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type UserServiceClient struct {
	Client pb.UserServiceClient
}

func NewUserClient(addr string) *UserServiceClient {
	conn, err := grpc.NewClient(addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("failed to connect to user service: %v", err)
	}

	return &UserServiceClient{
		Client: pb.NewUserServiceClient(conn),
	}
}

func (c *UserServiceClient) VerifyUserExists(userID uint32) (bool, error) {
	_, err := c.Client.GetUser(context.Background(), &pb.UserRequest{Id: userID})
	if err != nil {
		return false, err
	}
	return true, nil
}
