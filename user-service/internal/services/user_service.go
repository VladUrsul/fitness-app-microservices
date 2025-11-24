package services

import (
	"fitness-app-microservices/user-service/internal/db"
	"fitness-app-microservices/user-service/internal/domain/models"
)

func CreateUserService(u *models.User) (*models.User, error) {
	if err := db.DB.Create(u).Error; err != nil {
		return nil, err
	}
	return u, nil
}

func GetUserService(id uint) (*models.User, error) {
	var user models.User
	if err := db.DB.First(&user, id).Error; err != nil {
		return nil, err
	}
	return &user, nil
}
