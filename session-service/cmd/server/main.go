package main

import (
	"fitness-app-microservices/session-service/internal/db"
	"fitness-app-microservices/session-service/internal/handlers"

	docs "fitness-app-microservices/session-service/internal/docs"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func main() {
	db.Connect()

	docs.SwaggerInfo.BasePath = "/api/v1"

	r := gin.Default()
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	api := r.Group("/api/v1")
	api.POST("/sessions", handlers.CreateSession)
	api.GET("/sessions/:id", handlers.GetSession)

	r.Run(":8083")
}
