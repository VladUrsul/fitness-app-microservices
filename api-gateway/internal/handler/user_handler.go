package handler

import (
	"context"
	_ "fitness-app-microservices/api-gateway/internal/docs"
	pb "fitness-app-microservices/proto"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// GetUser godoc
// @Summary Get a user by ID
// @Description Retrieve a user from the UserService via gRPC
// @Tags Users
// @Produce json
// @Param id path int true "User ID"
// @Success 200 {object} UserResponseDoc
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Security BearerAuth
// @Router /users/{id} [get]
func (h *Handler) GetUser(c *gin.Context) {
	idStr := c.Param("id")

	idInt, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid user id"})
		return
	}

	resp, err := h.UserClient.GetUser(context.Background(), &pb.UserRequest{
		Id: uint32(idInt),
	})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, resp)
}

// CreateUser godoc
// @Summary Create a new user
// @Description Create a user via UserService gRPC
// @Tags Users
// @Accept json
// @Produce json
// @Param user body UserRequestDoc true "User payload"
// @Success 200 {object} UserResponseDoc
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Security BearerAuth
// @Router /users [post]
func (h *Handler) CreateUser(c *gin.Context) {
	var req pb.CreateUserRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	resp, err := h.UserClient.CreateUser(context.Background(), &req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, resp)
}
