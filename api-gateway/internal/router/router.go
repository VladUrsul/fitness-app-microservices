package router

import (
	"fitness-app-microservices/api-gateway/internal/config"
	_ "fitness-app-microservices/api-gateway/internal/docs"
	"fitness-app-microservices/api-gateway/internal/handler"
	"fitness-app-microservices/api-gateway/internal/middleware"

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	"github.com/gin-gonic/gin"
)

func SetupRouter(cfg *config.Config) *gin.Engine {
	r := gin.Default()

	// Middlewares
	r.Use(middleware.Logger())
	r.Use(middleware.Auth())

	// Handlers
	h := handler.NewHandler(cfg)
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	api := r.Group("/api")
	{
		api.GET("/users/:id", h.GetUser)
		api.GET("/sessions/:id", h.GetSession)
		api.GET("/workouts/:id", h.GetWorkout)

		api.POST("/users", h.CreateUser)
		api.POST("/sessions", h.CreateSession)
		api.POST("/workouts", h.CreateWorkout)
	}

	return r
}
