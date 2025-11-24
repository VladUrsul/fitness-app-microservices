package main

import (
	"fitness-app-microservices/api-gateway/internal/config"
	"fitness-app-microservices/api-gateway/internal/router"
	"log"
)

// @title Fitness App API Gateway
// @version 1.0
// @description API Gateway for fitness microservices
// @host localhost:8080
// @BasePath /api
func main() {
	cfg := config.LoadConfig()

	r := router.SetupRouter(cfg)

	log.Printf("Starting API Gateway on port %s", cfg.Port)
	if err := r.Run(":" + cfg.Port); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
