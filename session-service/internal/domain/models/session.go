package models

import "time"

type Session struct {
	ID         uint      `json:"id" gorm:"primaryKey;autoIncrement"`
	WorkoutID  uint      `json:"workout_id"`
	StartedAt  time.Time `json:"started_at"`
	FinishedAt time.Time `json:"finished_at"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
}
