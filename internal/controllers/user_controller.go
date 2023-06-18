package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/Karthika-Rajagopal/go-restful-backend/internal/models"
	"github.com/Karthika-Rajagopal/go-restful-backend/internal/repositories"
)

// UserController represents the user controller
type UserController struct {
	UserRepository repositories.UserRepository
}

// NewUserController creates a new instance of UserController
func NewUserController(userRepo repositories.UserRepository) *UserController {
	return &UserController{
		UserRepository: userRepo,
	}
}

// GetProfile handles the get profile API
func (uc *UserController) GetProfile(c *gin.Context) {
	userID := c.MustGet("userID").(uint)

	user, err := uc.UserRepository.GetUserByID(userID)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Profile not found"})
		return
	}

	c.JSON(http.StatusOK, user)
}

// UpdateProfile handles the update profile API
func (uc *UserController) UpdateProfile(c *gin.Context) {
	userID := c.MustGet("userID").(uint)

	var updateRequest models.UpdateProfileRequest
	if err := c.ShouldBindJSON(&updateRequest); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request payload"})
		return
	}

	// Perform Metamask conversion logic and update the user profile
	// ...

	// Update the user profile in the database
	if err := uc.UserRepository.UpdateUserProfile(userID, updateRequest.Address); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update profile"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Profile updated successfully"})
}
