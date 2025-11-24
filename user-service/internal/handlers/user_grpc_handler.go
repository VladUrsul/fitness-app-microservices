package handlers

import (
	"context"
	"errors"

	pb "fitness-app-microservices/proto"
	"fitness-app-microservices/user-service/internal/domain/models"
	"fitness-app-microservices/user-service/internal/services"

	"google.golang.org/protobuf/types/known/timestamppb"
	"gorm.io/gorm"
)

type UserServiceServer struct {
	pb.UnimplementedUserServiceServer
}

// CreateUser gRPC method
func (u *UserServiceServer) CreateUser(ctx context.Context, req *pb.CreateUserRequest) (*pb.UserResponse, error) {
	user := &models.User{
		Name:  req.Name,
		Email: req.Email,
	}

	created, err := services.CreateUserService(user)
	if err != nil {
		return nil, err
	}

	return mapToResponse(created), nil
}

// GetUser gRPC method
func (u *UserServiceServer) GetUser(ctx context.Context, req *pb.UserRequest) (*pb.UserResponse, error) {
	user, err := services.GetUserService(uint(req.Id))
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("user not found")
		}
		return nil, err
	}

	return mapToResponse(user), nil
}

// Mapper
func mapToResponse(u *models.User) *pb.UserResponse {
	return &pb.UserResponse{
		Id:        uint32(u.ID),
		Name:      u.Name,
		Email:     u.Email,
		CreatedAt: timestamppb.New(u.CreatedAt),
		UpdatedAt: timestamppb.New(u.UpdatedAt),
	}
}
