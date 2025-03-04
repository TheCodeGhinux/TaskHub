package routers

// import (
// 	"your_project/handlers"
// 	"your_project/middleware"

// 	"github.com/gin-gonic/gin"
// )

// func RegisterTaskRoutes(r *gin.RouterGroup) {
// 	taskGroup := r.Group("/tasks")

// 	// Only Admin & Manager can create tasks
// 	taskGroup.POST("/", middleware.RBAC("admin", "manager"), handlers.CreateTask)

// 	// All users within the tenant can view tasks
// 	taskGroup.GET("/", middleware.RBAC("admin", "manager", "user"), handlers.GetTasks)

// 	// Only Admin & Manager can update or delete tasks
// 	taskGroup.PUT("/:id", middleware.RBAC("admin", "manager"), handlers.UpdateTask)
// 	taskGroup.DELETE("/:id", middleware.RBAC("admin", "manager"), handlers.DeleteTask)
// }