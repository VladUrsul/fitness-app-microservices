package main

import (
	docs "fitness-app-microservices/workout-service/internal/docs"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	"fitness-app-microservices/workout-service/internal/db"
	"fitness-app-microservices/workout-service/internal/handlers"
)

func main() {
	db.Connect()
	docs.SwaggerInfo.BasePath = "/api/v1"

	r := gin.Default()
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	api := r.Group("/api/v1")
	api.POST("/workouts", handlers.CreateWorkout)
	api.GET("/workouts/:id", handlers.GetWorkout)

	r.Run(":8082")
}
