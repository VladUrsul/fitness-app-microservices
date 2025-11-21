package models

import "time"

type User struct {
	ID        uint      `json:"id" gorm:"primaryKey;autoIncrement"`
	Name      string    `json:"name" example:"Mihail"`
	Email     string    `json:"email" gorm:"uniqueIndex" example:"mihail.marinescu@example.com"`
	CreatedAt time.Time `json:"created_at" example:"2025-11-20T10:00:00Z"`
	UpdatedAt time.Time `json:"updated_at" example:"2025-11-20T10:00:00Z"`
}
