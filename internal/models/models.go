package models

import (
	"gorm.io/gorm"
)

// User represents the user entity
type User struct {
	gorm.Model
	//username string
	Email    string
	Password string
	// Add other fields as needed
}

// RegisterRequest represents the request payload for user registration
type RegisterRequest struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required"`
}

// LoginRequest represents the request payload for user login
type LoginRequest struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required"`
}

// UpdateProfileRequest represents the request payload for updating user profile
type UpdateProfileRequest struct {
	Address string `json:"address" binding:"required"`
}
