package handlers

import (
	"context"
	"errors"

	pb "fitness-app-microservices/proto"
	"fitness-app-microservices/session-service/internal/domain/models"
	"fitness-app-microservices/session-service/internal/services"

	"google.golang.org/protobuf/types/known/timestamppb"
	"gorm.io/gorm"
)

type SessionServiceServer struct {
	pb.UnimplementedSessionServiceServer
}

func (s *SessionServiceServer) CreateSession(ctx context.Context, req *pb.SessionRequest) (*pb.SessionResponse, error) {
	session := &models.Session{
		WorkoutID:  uint(req.WorkoutId),
		StartedAt:  req.StartedAt.AsTime(),
		FinishedAt: req.FinishedAt.AsTime(),
	}

	created, err := services.CreateSessionService(session)
	if err != nil {
		return nil, err
	}

	return mapToResponse(created), nil
}

func (s *SessionServiceServer) GetSession(ctx context.Context, req *pb.SessionRequest) (*pb.SessionResponse, error) {
	session, err := services.GetSessionService(uint(req.Id))
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("session not found")
		}
		return nil, err
	}

	return mapToResponse(session), nil
}

func mapToResponse(s *models.Session) *pb.SessionResponse {
	return &pb.SessionResponse{
		Id:         uint32(s.ID),
		WorkoutId:  uint32(s.WorkoutID),
		StartedAt:  timestamppb.New(s.StartedAt),
		FinishedAt: timestamppb.New(s.FinishedAt),
		CreatedAt:  timestamppb.New(s.CreatedAt),
		UpdatedAt:  timestamppb.New(s.UpdatedAt),
	}
}
