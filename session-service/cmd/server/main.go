package main

import (
	"log"
	"net"

	pb "fitness-app-microservices/proto"
	"fitness-app-microservices/session-service/internal/handlers"

	"google.golang.org/grpc"
)

func main() {
	lis, err := net.Listen("tcp", ":50052")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer()
	pb.RegisterSessionServiceServer(s, &handlers.SessionServiceServer{})

	log.Println("SessionService gRPC running on :50052")
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
