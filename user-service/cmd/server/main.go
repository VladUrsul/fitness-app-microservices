package main

import (
	docs "fitness-app-microservices/user-service/internal/docs"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	"fitness-app-microservices/user-service/internal/db"
	"fitness-app-microservices/user-service/internal/handlers"
)

// @title User Service API
// @version 1.0
// @description User microservice for fitness app
// @BasePath /api/v1
func main() {
	db.Connect()
	docs.SwaggerInfo.BasePath = "/api/v1"

	r := gin.Default()
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	api := r.Group("/api/v1")
	api.POST("/users", handlers.CreateUser)
	api.GET("/users/:id", handlers.GetUser)

	r.Run(":8081")
}
