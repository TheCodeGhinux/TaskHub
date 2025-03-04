package routing

import (
	"fmt"
	"log"

	"github.com/TheCodeGhinux/TaskHub/taskhub/pkg/config"
	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/swaggo/swag/example/basic/docs"
)

func Route() {

	configs := config.LoadConfig()
	r := gin.Default()
	// Swagger docs route
	docs.SwaggerInfo.BasePath = "api/v1"
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
	RouteRegister(r)

	err := r.Run(fmt.Sprintf("%s:%s", configs.Server.Host, configs.Server.Port))

	// **Logging registered routes**
	for _, route := range r.Routes() {
		log.Printf("Registered route: %s %s\n", route.Method, route.Path)
	}

	if err != nil {
		log.Fatal("Error starting server in routing: ", err)
		return
	}

}
