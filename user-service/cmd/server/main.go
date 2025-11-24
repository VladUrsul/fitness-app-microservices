package main

import (
	"log"
	"net"

	pb "fitness-app-microservices/proto"
	"fitness-app-microservices/user-service/internal/db"
	"fitness-app-microservices/user-service/internal/handlers"

	"google.golang.org/grpc"
)

func main() {
	db.Connect()

	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer()
	pb.RegisterUserServiceServer(s, &handlers.UserServiceServer{})

	log.Println("User gRPC service running on :50051")
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
