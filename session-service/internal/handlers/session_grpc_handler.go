package handlers

import (
	"context"
	"errors"

	pb "fitness-app-microservices/proto"
	"fitness-app-microservices/session-service/internal/clients"
	"fitness-app-microservices/session-service/internal/db"
	"fitness-app-microservices/session-service/internal/domain/models"

	"google.golang.org/protobuf/types/known/timestamppb"
	"gorm.io/gorm"
)

type SessionServiceServer struct {
	pb.UnimplementedSessionServiceServer
}

// CreateSession gRPC method
func (s *SessionServiceServer) CreateSession(ctx context.Context, req *pb.SessionRequest) (*pb.SessionResponse, error) {
	if !clients.VerifyWorkoutExists(uint(req.WorkoutId)) {
		return nil, errors.New("workout does not exist")
	}

	startedAt := req.StartedAt.AsTime()
	finishedAt := req.FinishedAt.AsTime()

	session := &models.Session{
		WorkoutID:  uint(req.WorkoutId),
		StartedAt:  startedAt,
		FinishedAt: finishedAt,
	}

	if err := db.DB.Create(session).Error; err != nil {
		return nil, err
	}

	return mapToResponse(session), nil
}

// GetSession gRPC method
func (s *SessionServiceServer) GetSession(ctx context.Context, req *pb.SessionRequest) (*pb.SessionResponse, error) {
	var session models.Session
	if err := db.DB.First(&session, req.Id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("session not found")
		}
		return nil, err
	}
	return mapToResponse(&session), nil
}

// Helper: map Session model to gRPC response
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
