package routers

import (
	"fmt"

	"github.com/TheCodeGhinux/TaskHub/taskhub/pkg/controllers/task"
	"github.com/TheCodeGhinux/TaskHub/taskhub/pkg/middlewares"
	"github.com/TheCodeGhinux/TaskHub/taskhub/pkg/repository/db"
	"github.com/TheCodeGhinux/TaskHub/taskhub/services/task"
	"github.com/gin-gonic/gin"
)

func TaskRoutes(router *gin.Engine, apiVersion string) *gin.Engine {
	taskService := services.NewTaskService(db.DB.Postgres)
	taskController := controllers.NewTaskController(taskService)

	taskGroup := router.Group(fmt.Sprintf("%v/tasks", apiVersion))
	taskGroup.Use(middlewares.JWTAuthMiddleware("your-secret-key"))

	{
		taskGroup.GET("/", taskController.GetAllTasks)                                               
		taskGroup.POST("/", middlewares.RBAC("admin", "manager"), taskController.CreateTask)
		taskGroup.PUT("/:id", middlewares.RBAC("admin", "manager"), taskController.UpdateTask)
		taskGroup.DELETE("/:id", middlewares.RBAC("admin", "manager"), taskController.DeleteTask)
	}

	return router
}
