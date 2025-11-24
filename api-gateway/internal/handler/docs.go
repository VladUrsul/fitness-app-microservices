package handler

import "time"

// SessionResponseDoc is used only for Swagger docs
type SessionResponseDoc struct {
	Id         uint32    `json:"id"`
	WorkoutId  uint32    `json:"workout_id"`
	StartedAt  time.Time `json:"started_at"`
	FinishedAt time.Time `json:"finished_at"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
}

// UserResponseDoc for Swagger
type UserResponseDoc struct {
	Id        uint32    `json:"id"`
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

// WorkoutResponseDoc for Swagger
type WorkoutResponseDoc struct {
	Id        uint32    `json:"id"`
	UserId    uint32    `json:"user_id"`
	Type      string    `json:"type"`
	Duration  int32     `json:"duration_minutes"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

// Request payloads (optional, for POST annotations)
type SessionRequestDoc struct {
	WorkoutId  uint32    `json:"workout_id"`
	StartedAt  time.Time `json:"started_at"`
	FinishedAt time.Time `json:"finished_at"`
}

type UserRequestDoc struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}

type WorkoutRequestDoc struct {
	UserId uint32 `json:"user_id"`
	Type   string `json:"type"`
}
