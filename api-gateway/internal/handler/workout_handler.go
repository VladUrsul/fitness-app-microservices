package handler

import (
	"context"
	"net/http"
	"strconv"

	pb "fitness-app-microservices/proto"

	"github.com/gin-gonic/gin"
)

func (h *Handler) GetWorkout(c *gin.Context) {
	idStr := c.Param("id")
	idInt, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid workout id"})
		return
	}
	id := uint32(idInt)

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
