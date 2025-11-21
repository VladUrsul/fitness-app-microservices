package services

import (
	"fitness-app-microservices/user-service/internal/db"
	"fitness-app-microservices/user-service/internal/domain/models"
)

func CreateUser(u *models.User) (*models.User, error) {
	if err := db.DB.Create(u).Error; err != nil {
		return nil, err
	}
	return u, nil
}

func GetUser(id uint) (*models.User, error) {
	var u models.User
	if err := db.DB.First(&u, id).Error; err != nil {
		return nil, err
	}
	return &u, nil
}
