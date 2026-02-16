package handler

import (
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/hadi-projects/go-react-starter/internal/dto"
	"github.com/hadi-projects/go-react-starter/internal/service"
)

type UserHandler interface {
	Register(c *gin.Context)
	Me(c *gin.Context)
}

type userHandler struct {
	service service.UserService
}

func NewUserHandler(service service.UserService) UserHandler {
	return &userHandler{service: service}
}

func (h *userHandler) Register(c *gin.Context) {
	var req dto.RegisterRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	res, err := h.service.Register(req)
	if err != nil {
		log.Printf("Register failed: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"data": res,
	})
}

func (h *userHandler) Me(c *gin.Context) {
	// userIDStr := c.GetString("user_id") // Assuming middleware sets this
	// If middleware sets it as a number:
	// userID := c.GetUint("user_id")

	// For now let's assume it comes from claims as float64 (jwt default) or string
	// Let's rely on middleware parsing it to proper type or just get it from context if set by AuthMiddleware

	// Temporary implementation assuming AuthMiddleware sets "userID" context variable
	val, exists := c.Get("userID")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}

	var userID uint
	switch v := val.(type) {
	case uint:
		userID = v
	case float64:
		userID = uint(v)
	case string:
		id, _ := strconv.ParseUint(v, 10, 32)
		userID = uint(id)
	}

	res, err := h.service.GetMe(userID)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": res,
	})
}
