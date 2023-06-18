package models

import "gorm.io/gorm"

// User represents the user model
type User struct {
	gorm.Model
	Email    string `gorm:"unique"`
	Password string
}

// RegisterRequest represents the request payload for user registration
type RegisterRequest struct {
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

// LoginRequest represents the request payload for user login
type LoginRequest struct {
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

// UpdateProfileRequest represents the request payload for updating user profile
type UpdateProfileRequest struct {
	Address string `json:"address" binding:"required"`
}
