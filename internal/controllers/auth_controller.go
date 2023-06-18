package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"Karthika-Rajagopal/go-restful-backend/internal/models"
	"Karthika-Rajagopal/go-restful-backend/internal/repositories"
	"Karthika-Rajagopal/go-restful-backend/internal/utils"
)

// AuthController represents the authentication controller
type AuthController struct {
	UserRepository repositories.UserRepository
}

// NewAuthController creates a new instance of AuthController
func NewAuthController(userRepo repositories.UserRepository) *AuthController {
	return &AuthController{
		UserRepository: userRepo,
	}
}

// Register handles the user registration API
func (ac *AuthController) Register(c *gin.Context) {
	var registerRequest models.RegisterRequest
	if err := c.ShouldBindJSON(&registerRequest); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request payload"})
		return
	}

	user := models.User{
		Email:    registerRequest.Email,
		Password: registerRequest.Password,
	}

	if err := ac.UserRepository.CreateUser(&user); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create user"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "User created successfully"})
}

// Login handles the user login API
func (ac *AuthController) Login(c *gin.Context) {
	var loginRequest models.LoginRequest
	if err := c.ShouldBindJSON(&loginRequest); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request payload"})
		return
	}

	user, err := ac.UserRepository.GetUserByEmail(loginRequest.Email)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid credentials"})
		return
	}

	if user.Password != loginRequest.Password {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid credentials"})
		return
	}

	token, err := utils.GenerateJWTToken(user.ID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate token"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"token": token, "user": user})
}
