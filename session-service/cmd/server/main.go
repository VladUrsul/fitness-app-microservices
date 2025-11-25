package main

import (
	"log"
	"net"

	pb "fitness-app-microservices/proto"
	"fitness-app-microservices/session-service/internal/db"
	grpcClient "fitness-app-microservices/session-service/internal/grpc"
	"fitness-app-microservices/session-service/internal/handlers"

	"google.golang.org/grpc"
)

func main() {

	db.Connect()

	workoutClient := grpcClient.NewWorkoutClient("workout-service:50052")

	lis, err := net.Listen("tcp", ":50053")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer()
	pb.RegisterSessionServiceServer(s, &handlers.SessionServiceServer{
		WorkoutClient: workoutClient,
	})

	log.Println("SessionService gRPC running on :50053")
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
