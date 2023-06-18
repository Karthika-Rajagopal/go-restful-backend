package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"Karthika-Rajagopal/go-restful-backend/internal/models"
	"Karthika-Rajagopal/go-restful-backend/internal/repositories"
	//"Karthika-Rajagopal/go-restful-backend/internal/utils"
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

// GetProfile handles the user profile retrieval API
func (uc *UserController) GetProfile(c *gin.Context) {
	userID := c.MustGet("userID").(uint)
	user, err := uc.UserRepository.GetUserByID(userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve user profile"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"user": user})
}

// UpdateProfile handles the user profile update API
func (uc *UserController) UpdateProfile(c *gin.Context) {
	userID := c.MustGet("userID").(uint)

	var updateProfileRequest models.UpdateProfileRequest
	if err := c.ShouldBindJSON(&updateProfileRequest); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request payload"})
		return
	}

	// Perform the required logic to convert the signed string to public address
	// and update it in the database

	if err := uc.UserRepository.UpdateUserProfile(userID, updateProfileRequest.Address); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update user profile"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "User profile updated successfully"})
}
