package handler

import (
	"context"
	"net/http"

	pb "fitness-app-microservices/api-gateway/proto"

	"github.com/gin-gonic/gin"
)

func (h *Handler) GetWorkout(c *gin.Context) {
	id := c.Param("id")
	resp, err := h.WorkoutClient.GetWorkout(context.Background(), &pb.WorkoutRequest{Id: id})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, resp)
}
