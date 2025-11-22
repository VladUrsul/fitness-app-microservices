package handler

import (
	"context"
	"net/http"

	pb "fitness-app-microservices/proto"

	"github.com/gin-gonic/gin"
)

func (h *Handler) GetWorkout(c *gin.Context) {
	id := c.Param("id")
	resp, err := h.WorkoutClient.GetWorkout(context.Background(), &pb.GetWorkoutRequest{Id: id})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, resp)
}

func (h *Handler) CreateWorkout(c *gin.Context) {
	var req pb.CreateWorkoutRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	resp, err := h.WorkoutClient.CreateWorkout(context.Background(), &req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, resp)
}
