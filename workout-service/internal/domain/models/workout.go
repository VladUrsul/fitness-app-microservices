package models

import "time"

type Workout struct {
	ID        uint      `json:"id" gorm:"primaryKey;autoIncrement"`
	UserID    uint      `json:"user_id"`
	Title     string    `json:"title"`
	Duration  int32     `json:"duration_minutes"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
