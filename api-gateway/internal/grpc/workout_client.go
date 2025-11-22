package grpc

import (
	"log"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	pb "fitness-app-microservices/api-gateway/proto"
)

func NewWorkoutClient(addr string) pb.WorkoutServiceClient {
	conn, err := grpc.NewClient(addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("Failed to connect to workout service: %v", err)
	}
	return pb.NewWorkoutServiceClient(conn)
}
