package main

import (
	"log"
	"net"

	pb "fitness-app-microservices/proto"
	"fitness-app-microservices/workout-service/internal/db"
	grpcClient "fitness-app-microservices/workout-service/internal/grpc"
	"fitness-app-microservices/workout-service/internal/handlers"

	"google.golang.org/grpc"
)

func main() {
	db.Connect()

	userClient := grpcClient.NewUserClient("user-service:50051")

	lis, err := net.Listen("tcp", ":50052")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer()
	pb.RegisterWorkoutServiceServer(s, &handlers.WorkoutServiceServer{
		UserClient: userClient,
	})
	log.Println("WorkoutService gRPC running on :50052")
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
