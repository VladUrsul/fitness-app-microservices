package services

import (
	"errors"

	"fitness-app-microservices/workout-service/internal/db"
	"fitness-app-microservices/workout-service/internal/domain/models"
	"fitness-app-microservices/workout-service/internal/grpc"
)

func CreateWorkoutService(w *models.Workout, userClient *grpc.UserServiceClient) (*models.Workout, error) {
	exists, err := userClient.VerifyUserExists(uint32(w.UserID))
	if err != nil || !exists {
		return nil, errors.New("user does not exist")
	}

	if err := db.DB.Create(w).Error; err != nil {
		return nil, err
	}
	return w, nil
}

func GetWorkoutService(id uint) (*models.Workout, error) {
	var w models.Workout
	if err := db.DB.First(&w, id).Error; err != nil {
		return nil, err
	}
	return &w, nil
}
