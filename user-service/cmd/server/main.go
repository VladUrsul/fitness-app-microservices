package main

import (
	"log"
	"net"

	"fitness-app-microservices/user-service/internal/db"
	grpcServer "fitness-app-microservices/user-service/internal/grpc"

	"google.golang.org/grpc"
)

func main() {
	db.Connect()

	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer()
	grpcServer.RegisterUserServiceServer(s)

	log.Println("User gRPC service running on :50051")
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
