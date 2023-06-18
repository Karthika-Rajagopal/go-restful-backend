package repositories

import (
	"errors"

	"Karthika-Rajagopal/go-restful-backend/internal/models"
	"gorm.io/gorm"
)

// UserRepository represents the user repository
type UserRepository struct {
	DB *gorm.DB
}

// NewUserRepository creates a new instance of UserRepository
func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{
		DB: db,
	}
}

// CreateUser creates a new user in the database
func (ur *UserRepository) CreateUser(user *models.User) error {
	result := ur.DB.Create(user)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

// GetUserByEmail retrieves a user by email from the database
func (ur *UserRepository) GetUserByEmail(email string) (*models.User, error) {
	var user models.User
	result := ur.DB.Where("email = ?", email).First(&user)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, errors.New("user not found")
		}
		return nil, result.Error
	}
	return &user, nil
}

// GetUserByID retrieves a user by ID from the database
func (ur *UserRepository) GetUserByID(userID uint) (*models.User, error) {
	var user models.User
	result := ur.DB.First(&user, userID)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, errors.New("user not found")
		}
		return nil, result.Error
	}
	return &user, nil
}

// UpdateUserProfile updates the user profile in the database
func (ur *UserRepository) UpdateUserProfile(userID uint, address string) error {
	result := ur.DB.Model(&models.User{}).Where("id = ?", userID).Update("address", address)
	if result.Error != nil {
		return result.Error
	}
	return nil
}
