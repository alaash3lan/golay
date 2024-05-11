package handler

import (
	"golay/internal/domain/user/model"
	"golay/internal/domain/user/service"
	"golay/internal/domain/user/validators"
	"golay/internal/utils"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// UserHandler handles HTTP requests related to users.
type UserHandler struct {
	service  *service.UserService
}

// NewUserHandler creates a new UserHandler instance.
func NewUserHandler(service *service.UserService) *UserHandler {
	return &UserHandler{service: service}
}

// CreateUser handles the creation of a new user.
func (h *UserHandler) CreateUser(c *gin.Context) {

	
	// var newUser model.User
	var request validators.CreateUserRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

		// Validate the request
		validationErrors := request.Validate()
		if validationErrors != nil {
			c.JSON(http.StatusBadRequest, gin.H{"errors": validationErrors})
			return
		}
	

	newUser := request.ToUserModel()

	if err := h.service.CreateUser(newUser); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, newUser)
}

func (h *UserHandler) GetUser(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID"})
		return
	}
	user, err := h.service.GetUser(uint(id))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}
	c.JSON(http.StatusOK, user)
}

// GetUser handles retrieving a user by ID.
func (h *UserHandler) ListUsers(c *gin.Context) {
	user, err := h.service.ListUsers()
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}
	utils.Respond(c, utils.Response{
		Status:  http.StatusOK,
		Message: "success",
		Data:    user,
	})
}

// UpdateUser handles updating an existing user.
func (h *UserHandler) UpdateUser(c *gin.Context) {
	var updatedUser model.User
	if err := c.ShouldBindJSON(&updatedUser); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := h.service.UpdateUser(&updatedUser); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, updatedUser)
}

// DeleteUser handles deleting a user by ID.
func (h *UserHandler) DeleteUser(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID"})
		return
	}
	if err := h.service.DeleteUser(uint(id)); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "User deleted"})
}
