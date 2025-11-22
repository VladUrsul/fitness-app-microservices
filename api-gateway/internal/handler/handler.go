package handler

import (
	"fitness-app-microservices/api-gateway/internal/config"
	"fitness-app-microservices/api-gateway/internal/grpc"
	pb "fitness-app-microservices/api-gateway/proto"
)

type Handler struct {
	UserClient    pb.UserServiceClient
	SessionClient pb.SessionServiceClient
	WorkoutClient pb.WorkoutServiceClient
}

func NewHandler(cfg *config.Config) *Handler {
	return &Handler{
		UserClient:    grpc.NewUserClient(cfg.UserSvcAddr),
		SessionClient: grpc.NewSessionClient(cfg.SessionSvcAddr),
		WorkoutClient: grpc.NewWorkoutClient(cfg.WorkoutSvcAddr),
	}
}
