package services

import (
	"errors"
	"fitness-app-microservices/session-service/internal/clients"
	"fitness-app-microservices/session-service/internal/db"
	"fitness-app-microservices/session-service/internal/domain/models"
)

func CreateSessionService(s *models.Session) (*models.Session, error) {
	if !clients.VerifyWorkoutExists(s.WorkoutID) {
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
