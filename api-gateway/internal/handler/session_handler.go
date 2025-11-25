package handler

import (
	"context"
	_ "fitness-app-microservices/api-gateway/internal/docs"
	"fitness-app-microservices/api-gateway/internal/dto"
	pb "fitness-app-microservices/proto"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"google.golang.org/protobuf/types/known/timestamppb"
)

// GetSession godoc
// @Summary Get a session by ID
// @Description Retrieve a session from the SessionService via gRPC
// @Tags Sessions
// @Produce json
// @Param id path int true "Session ID"
// @Success 200 {object} SessionResponseDoc
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Security BearerAuth
// @Router /sessions/{id} [get]
func (h *Handler) GetSession(c *gin.Context) {
	idStr := c.Param("id")

	// Convert string to uint32
	idInt, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid session id"})
		return
	}
	id := uint32(idInt)

	// Call gRPC
	resp, err := h.SessionClient.GetSession(context.Background(), &pb.SessionRequest{Id: id})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, resp)
}

// CreateSession godoc
// @Summary Create a new session
// @Description Create a session via SessionService gRPC
// @Tags Sessions
// @Accept json
// @Produce json
// @Param session body SessionRequestDoc true "Session payload"
// @Success 200 {object} SessionResponseDoc
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Security BearerAuth
// @Router /sessions [post]
func (h *Handler) CreateSession(c *gin.Context) {
	var req dto.CreateSessionRequestHTTP

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid json"})
		return
	}

	start, err := time.Parse(time.RFC3339, req.StartedAt)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid started_at"})
		return
	}

	finish, err := time.Parse(time.RFC3339, req.FinishedAt)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid finished_at"})
		return
	}

	grpcReq := &pb.SessionRequest{
		WorkoutId:  req.WorkoutId,
		StartedAt:  timestamppb.New(start),
		FinishedAt: timestamppb.New(finish),
	}

	resp, err := h.SessionClient.CreateSession(context.Background(), grpcReq)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, resp)
}
