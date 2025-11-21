package models

import "time"

type Workout struct {
	ID        uint      `json:"id" gorm:"primaryKey;autoIncrement"`
	UserID    uint      `json:"user_id" example:"1"`
	Type      string    `json:"type" example:"Yoga"`
	Scheduled time.Time `json:"scheduled"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
