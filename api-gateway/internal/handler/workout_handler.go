package handler

import (
	"context"
	_ "fitness-app-microservices/api-gateway/internal/docs"
	pb "fitness-app-microservices/proto"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// GetWorkout godoc
// @Summary Get a workout by ID
// @Description Retrieve a workout from the WorkoutService via gRPC
// @Tags Workouts
// @Produce json
// @Param id path int true "Workout ID"
// @Success 200 {object} WorkoutResponseDoc
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /workouts/{id} [get]
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

// CreateWorkout godoc
// @Summary Create a new workout
// @Description Create a workout via WorkoutService gRPC
// @Tags Workouts
// @Accept json
// @Produce json
// @Param workout body WorkoutRequestDoc true "Workout payload"
// @Success 200 {object} WorkoutResponseDoc
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /workouts [post]
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
