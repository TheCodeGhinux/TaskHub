package controllers

import (
	"net/http"

	"github.com/TheCodeGhinux/TaskHub/taskhub/internal/models"
	services "github.com/TheCodeGhinux/TaskHub/taskhub/services/task"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type TaskController struct {
	TaskService *services.TaskService
}

func NewTaskController(taskService *services.TaskService) *TaskController {
	return &TaskController{TaskService: taskService}
}

// GetAllTasks retrieves all tasks for the user's tenant.
func (tc *TaskController) GetAllTasks(c *gin.Context) {
	tenantID := c.GetString("tenant_id")
	tasks, err := tc.TaskService.GetAllTasks(tenantID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch tasks"})
		return
	}
	c.JSON(http.StatusOK, tasks)
}

// CreateTask creates a new task.
func (tc *TaskController) CreateTask(c *gin.Context) {
	var task models.Task
	if err := c.ShouldBindJSON(&task); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request payload"})
		return
	}

	tenantID, err := uuid.Parse(c.GetString("tenant_id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid tenant ID"})
		return
	}

	task.TenantID = tenantID
	if err := tc.TaskService.CreateTask(&task); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create task"})
		return
	}

	c.JSON(http.StatusCreated, task)
}

// UpdateTask updates an existing task.
func (tc *TaskController) UpdateTask(c *gin.Context) {
	taskID := c.Param("id")
	tenantID := c.GetString("tenant_id")

	var updates map[string]interface{}
	if err := c.ShouldBindJSON(&updates); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request payload"})
		return
	}

	if err := tc.TaskService.UpdateTask(taskID, tenantID, updates); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update task"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Task updated successfully"})
}

// DeleteTask deletes a task.
func (tc *TaskController) DeleteTask(c *gin.Context) {
	taskID := c.Param("id")
	tenantID := c.GetString("tenant_id")

	if err := tc.TaskService.DeleteTask(taskID, tenantID); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete task"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Task deleted successfully"})
}
