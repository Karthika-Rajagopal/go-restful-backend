package repositories

import (
	"errors"

	"gorm.io/gorm"

	"github.com/Karthika-Rajagopal/go-restful-backend/internal/models"
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

// GetUserByEmail retrieves a user from the database by email
func (ur *UserRepository) GetUserByEmail(email string) (*models.User, error) {
	var user models.User
	result := ur.DB.Where("email = ?", email).First(&user)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return nil, result.Error
	}
	return &user, nil
}

// GetUserByID retrieves a user from the database by ID
func (ur *UserRepository) GetUserByID(id uint) (*models.User, error) {
	var user models.User
	result := ur.DB.First(&user, id)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return nil, result.Error
	}
	return &user, nil
}

// UpdateUserProfile updates the user profile in the database
func (ur *UserRepository) UpdateUserProfile(id uint, address string) error {
	result := ur.DB.Model(&models.User{}).Where("id = ?", id).Updates(models.User{Address: address})
	if result.Error != nil {
		return result.Error
	}
	return nil
}
