package routing

import (
	"github.com/TheCodeGhinux/TaskHub/taskhub/pkg/routers"
	"github.com/gin-gonic/gin"
)

func RouteRegister(router *gin.Engine) {

	apiVersion := "api/v1"
	routers.Greeting(router, apiVersion)
	routers.AuthRoutes(router, apiVersion)

}
