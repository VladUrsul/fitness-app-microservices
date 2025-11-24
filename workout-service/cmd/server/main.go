package main

import (
	"log"
	"net"

	pb "fitness-app-microservices/proto"
	"fitness-app-microservices/workout-service/internal/db"
	"fitness-app-microservices/workout-service/internal/handlers"

	"google.golang.org/grpc"
)

func main() {
	db.Connect()

	lis, err := net.Listen("tcp", ":50053") // assign port
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer()
	pb.RegisterWorkoutServiceServer(s, &handlers.WorkoutServiceServer{})

	log.Println("WorkoutService gRPC running on :50053")
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
