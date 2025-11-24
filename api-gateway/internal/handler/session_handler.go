package handler

import (
	"context"
	_ "fitness-app-microservices/api-gateway/internal/docs"
	pb "fitness-app-microservices/proto"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
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
// @Router /sessions [post]
func (h *Handler) CreateSession(c *gin.Context) {
	var req pb.SessionRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	resp, err := h.SessionClient.CreateSession(context.Background(), &req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, resp)
}
