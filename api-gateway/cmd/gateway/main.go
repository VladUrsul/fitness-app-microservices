package main

import (
	"fitness-app-microservices/api-gateway/internal/config"
	"fitness-app-microservices/api-gateway/internal/router"
	"log"
)

func main() {
	cfg := config.LoadConfig()

	r := router.SetupRouter(cfg)

	log.Printf("Starting API Gateway on port %s", cfg.Port)
	if err := r.Run(":" + cfg.Port); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
