package handler

import (
	"context"
	"net/http"

	pb "fitness-app-microservices/api-gateway/proto"

	"github.com/gin-gonic/gin"
)

func (h *Handler) GetSession(c *gin.Context) {
	id := c.Param("id")
	resp, err := h.SessionClient.GetSession(context.Background(), &pb.SessionRequest{Id: id})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, resp)
}
