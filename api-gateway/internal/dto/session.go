package dto

type CreateSessionRequestHTTP struct {
	WorkoutId  uint32 `json:"workout_id"`
	StartedAt  string `json:"started_at"`
	FinishedAt string `json:"finished_at"`
}
