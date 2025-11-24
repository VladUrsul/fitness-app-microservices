package services

import (
	"errors"

	"fitness-app-microservices/session-service/internal/db"
	"fitness-app-microservices/session-service/internal/domain/models"
	grpcClient "fitness-app-microservices/session-service/internal/grpc"
)

func CreateSessionService(s *models.Session, workoutClient *grpcClient.WorkoutServiceClient) (*models.Session, error) {
	exists, err := workoutClient.VerifyWorkoutExists(uint32(s.WorkoutID))
	if err != nil || !exists {
		return nil, errors.New("workout does not exist")
	}

	if err := db.DB.Create(s).Error; err != nil {
		return nil, err
	}

	return s, nil
}

func GetSessionService(id uint) (*models.Session, error) {
	var s models.Session
	if err := db.DB.First(&s, id).Error; err != nil {
		return nil, err
	}
	return &s, nil
}
