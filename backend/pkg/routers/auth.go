package routers

import (
	"fmt"

	"github.com/gin-gonic/gin"

	controllers "github.com/TheCodeGhinux/TaskHub/taskhub/pkg/controllers/auth"
	"github.com/TheCodeGhinux/TaskHub/taskhub/pkg/repository/db"
	services "github.com/TheCodeGhinux/TaskHub/taskhub/services/auth"
	UserServices "github.com/TheCodeGhinux/TaskHub/taskhub/services/user"
)

// AuthRoutes sets up authentication endpoints.
func AuthRoutes(router *gin.Engine, ApiVersion string) *gin.Engine {
	userService := UserServices.NewUserService(db.DB.Postgres)
	authService := services.NewAuthService(userService)
	authController := controllers.NewAuthController(authService)

	AuthGroup := router.Group(fmt.Sprintf("%v/auth", ApiVersion))
	{
		AuthGroup.POST("/register", authController.Register)
		AuthGroup.POST("/login", authController.Login)
	}

	return router
}
