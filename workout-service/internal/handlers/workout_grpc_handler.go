package handlers

import (
	"context"
	"errors"

	pb "fitness-app-microservices/proto"
	"fitness-app-microservices/workout-service/internal/db"
	"fitness-app-microservices/workout-service/internal/domain/models"

	"gorm.io/gorm"
)

type WorkoutServiceServer struct {
	pb.UnimplementedWorkoutServiceServer
}

func (s *WorkoutServiceServer) CreateWorkout(ctx context.Context, req *pb.CreateWorkoutRequest) (*pb.WorkoutResponse, error) {
	w := &models.Workout{
		Title:    req.Title,
		Duration: req.DurationMinutes,
		UserID:   uint(req.UserId),
	}

	if err := db.DB.Create(w).Error; err != nil {
		return nil, err
	}

	return mapToResponse(w), nil
}

func (s *WorkoutServiceServer) GetWorkout(ctx context.Context, req *pb.GetWorkoutRequest) (*pb.WorkoutResponse, error) {
	var w models.Workout
	if err := db.DB.First(&w, req.Id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("workout not found")
		}
		return nil, err
	}

	return mapToResponse(&w), nil
}

func (s *WorkoutServiceServer) ListWorkouts(ctx context.Context, _ *pb.Empty) (*pb.WorkoutListResponse, error) {
	var workouts []models.Workout
	if err := db.DB.Find(&workouts).Error; err != nil {
		return nil, err
	}

	resp := &pb.WorkoutListResponse{}
	for _, w := range workouts {
		resp.Workouts = append(resp.Workouts, mapToResponse(&w))
	}

	return resp, nil
}

func mapToResponse(w *models.Workout) *pb.WorkoutResponse {
	return &pb.WorkoutResponse{
		Id:              uint32(w.ID),
		Title:           w.Title,
		DurationMinutes: w.Duration,
		UserId:          uint32(w.UserID),
	}
}
