package services

import (
	"errors"

	"github.com/google/uuid"
	"gorm.io/gorm"

	"github.com/TheCodeGhinux/TaskHub/taskhub/internal/models"
)

// UserService handles user-related operations.
type UserService struct {
	DB *gorm.DB
}

// NewUserService creates a new instance of UserService.
func NewUserService(db *gorm.DB) *UserService {
	return &UserService{DB: db}
}

// CreateUser inserts a new user into the database.
func (s *UserService) CreateUser(user *models.User) error {
	user.ID = uuid.New() // Ensure UUID is generated
	if err := s.DB.Create(user).Error; err != nil {
		return err
	}
	return nil
}

// GetUserByID retrieves a user by their ID.
func (s *UserService) GetUserByID(id uuid.UUID) (*models.User, error) {
	var user models.User
	if err := s.DB.Preload("Tenant").Where("id = ?", id).First(&user).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return &user, nil
}

// GetUserByEmail retrieves a user by their email.
func (s *UserService) GetUserByEmail(email string) (*models.User, error) {
	var user models.User
	if err := s.DB.Preload("Tenant").Where("email = ?", email).First(&user).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return &user, nil
}
