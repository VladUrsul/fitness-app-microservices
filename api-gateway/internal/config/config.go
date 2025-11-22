package config

import "os"

type Config struct {
	Port           string
	UserSvcAddr    string
	SessionSvcAddr string
	WorkoutSvcAddr string
}

func LoadConfig() *Config {
	return &Config{
		Port:           getEnv("GATEWAY_PORT", "8080"),
		UserSvcAddr:    getEnv("USER_SERVICE_ADDR", "localhost:50051"),
		SessionSvcAddr: getEnv("SESSION_SERVICE_ADDR", "localhost:50052"),
		WorkoutSvcAddr: getEnv("WORKOUT_SERVICE_ADDR", "localhost:50053"),
	}
}

func getEnv(key, fallback string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return fallback
}
