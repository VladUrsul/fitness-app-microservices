package grpc

import (
	"context"

	pb "fitness-app-microservices/proto"
	"fitness-app-microservices/user-service/internal/domain/models"
	"fitness-app-microservices/user-service/internal/services"

	"google.golang.org/grpc"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type UserServiceServer struct {
	pb.UnimplementedUserServiceServer
}

func RegisterUserServiceServer(s *grpc.Server) {
	pb.RegisterUserServiceServer(s, &UserServiceServer{})
}

func (u *UserServiceServer) CreateUser(ctx context.Context, req *pb.CreateUserRequest) (*pb.UserResponse, error) {
	user := &models.User{
		Name:  req.Name,
		Email: req.Email,
	}

	created, err := services.CreateUser(user)
	if err != nil {
		return nil, err
	}

	return &pb.UserResponse{
		Id:        uint32(created.ID),
		Name:      created.Name,
		Email:     created.Email,
		CreatedAt: timestamppb.New(created.CreatedAt),
		UpdatedAt: timestamppb.New(created.UpdatedAt),
	}, nil
}

func (u *UserServiceServer) GetUser(ctx context.Context, req *pb.UserRequest) (*pb.UserResponse, error) {
	user, err := services.GetUser(uint(req.Id))
	if err != nil {
		return nil, err
	}

	return &pb.UserResponse{
		Id:        uint32(user.ID),
		Name:      user.Name,
		Email:     user.Email,
		CreatedAt: timestamppb.New(user.CreatedAt),
		UpdatedAt: timestamppb.New(user.UpdatedAt),
	}, nil
}
