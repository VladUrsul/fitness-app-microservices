package grpc

import (
	"context"
	"log"

	pb "fitness-app-microservices/proto"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type WorkoutServiceClient struct {
	Client pb.WorkoutServiceClient
}

func NewWorkoutClient(addr string) *WorkoutServiceClient {
	conn, err := grpc.Dial(addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("failed to connect to workout service: %v", err)
	}

	return &WorkoutServiceClient{
		Client: pb.NewWorkoutServiceClient(conn),
	}
}

func (c *WorkoutServiceClient) VerifyWorkoutExists(workoutID uint32) (bool, error) {
	_, err := c.Client.GetWorkout(context.Background(), &pb.GetWorkoutRequest{Id: workoutID})
	if err != nil {
		return false, err
	}
	return true, nil
}
