package handler

import (
	"context"
	"net/http"

	pb "fitness-app-microservices/api-gateway/proto"

	"github.com/gin-gonic/gin"
)

func (h *Handler) GetUser(c *gin.Context) {
	id := c.Param("id")
	resp, err := h.UserClient.GetUser(context.Background(), &pb.UserRequest{Id: id})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, resp)
}
