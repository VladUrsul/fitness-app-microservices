package grpc

import (
	"context"
	"errors"
	"strconv"

	pb "fitness-app-microservices/proto"
	"fitness-app-microservices/session-service/internal/domain/models"
	"fitness-app-microservices/session-service/internal/services"

	"google.golang.org/protobuf/types/known/timestamppb"
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

	return &pb.SessionResponse{
		Id:         strconv.Itoa(int(created.ID)),
		WorkoutId:  uint32(created.WorkoutID),
		StartedAt:  timestamppb.New(created.StartedAt),
		FinishedAt: timestamppb.New(created.FinishedAt),
	}, nil
}

func (s *SessionServiceServer) GetSession(ctx context.Context, req *pb.SessionRequest) (*pb.SessionResponse, error) {
	id, err := strconv.Atoi(req.Id)
	if err != nil {
		return nil, errors.New("invalid session id")
	}

	session, err := services.GetSessionService(uint(id))
	if err != nil {
		return nil, err
	}

	return &pb.SessionResponse{
		Id:         strconv.Itoa(int(session.ID)),
		WorkoutId:  uint32(session.WorkoutID),
		StartedAt:  timestamppb.New(session.StartedAt),
		FinishedAt: timestamppb.New(session.FinishedAt),
	}, nil
}
