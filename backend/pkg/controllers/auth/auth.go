package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"

	services "github.com/TheCodeGhinux/TaskHub/taskhub/services/auth"
)

// AuthController handles authentication endpoints.
type AuthController struct {
	AuthService *services.AuthService
}

// NewAuthController creates a new instance of AuthController.
func NewAuthController(authService *services.AuthService) *AuthController {
	return &AuthController{AuthService: authService}
}

// Register endpoint
func (ac *AuthController) Register(c *gin.Context) {
	var userInput struct {
		Name     string `json:"name"`
		Password string `json:"password"`
	}

	if err := c.ShouldBindJSON(&userInput); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	// Pass input to AuthService
	err := ac.AuthService.RegisterUser(userInput.Name, userInput.Password)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "User registered successfully"})
}

// Login endpoint
func (ac *AuthController) Login(c *gin.Context) {
	var credentials struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	if err := c.ShouldBindJSON(&credentials); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	token, err := ac.AuthService.LoginUser(credentials.Email, credentials.Password)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"token": token})
}
