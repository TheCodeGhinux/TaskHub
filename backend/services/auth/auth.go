package services

import (
	"errors"
	"time"

	"github.com/dgrijalva/jwt-go"
	"golang.org/x/crypto/bcrypt"

	"github.com/TheCodeGhinux/TaskHub/taskhub/internal/models"
	userService "github.com/TheCodeGhinux/TaskHub/taskhub/services/user"
)

// JWT secret (should be stored in ENV variables)
var jwtSecret = []byte("your-secret-key")

// AuthService struct
type AuthService struct {
	UserService *userService.UserService
}

// NewAuthService constructor
func NewAuthService(userService *userService.UserService) *AuthService {
	return &AuthService{UserService: userService}
}

// Register a new user
func (s *AuthService) RegisterUser(name, password string) error {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return errors.New("failed to hash password")
	}

	user := models.User{
		Name:     name,
		Password: string(hashedPassword),
	}

	return s.UserService.CreateUser(&user)
}

// Login user
func (s *AuthService) LoginUser(email, password string) (string, error) {
	user, err := s.UserService.GetUserByEmail(email)
	if err != nil || user == nil {
		return "", errors.New("invalid email or password")
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
		return "", errors.New("invalid email or password")
	}

	// Generate JWT token
	claims := jwt.MapClaims{
		"user_id": user.ID.String(),
		"email":   user.Email,
		"exp":     time.Now().Add(time.Hour * 24).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(jwtSecret)
}
