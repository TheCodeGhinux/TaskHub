package routers

import (
	"fmt"

	controller "github.com/TheCodeGhinux/TaskHub/taskhub/pkg/controllers/greeting"
	"github.com/gin-gonic/gin"
)

func Greeting(router *gin.Engine, ApiVersion string) *gin.Engine {
	greetController := controller.Greet

	greetGroup := router.Group(fmt.Sprintf("%v", ApiVersion))
	{
		greetGroup.GET("/", greetController)
	}

	return router
}
