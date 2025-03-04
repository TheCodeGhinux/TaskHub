package services

import (
	"github.com/TheCodeGhinux/TaskHub/taskhub/internal/models"
	"gorm.io/gorm"
)

type TaskService struct {
	DB *gorm.DB
}

func NewTaskService(db *gorm.DB) *TaskService {
	return &TaskService{DB: db}
}

// GetAllTasks retrieves tasks for the authenticated user's tenant.
func (s *TaskService) GetAllTasks(tenantID string) ([]models.Task, error) {
	var tasks []models.Task
	if err := s.DB.Where("tenant_id = ?", tenantID).Find(&tasks).Error; err != nil {
		return nil, err
	}
	return tasks, nil
}

// CreateTask creates a new task for the authenticated user's tenant.
func (s *TaskService) CreateTask(task *models.Task) error {
	return s.DB.Create(task).Error
}

// UpdateTask updates an existing task while ensuring tenant isolation.
func (s *TaskService) UpdateTask(taskID string, tenantID string, updates map[string]interface{}) error {
	result := s.DB.Model(&models.Task{}).Where("id = ? AND tenant_id = ?", taskID, tenantID).Updates(updates)
	if result.RowsAffected == 0 {
		return gorm.ErrRecordNotFound
	}
	return result.Error
}

// DeleteTask ensures only tasks belonging to the tenant can be deleted.
func (s *TaskService) DeleteTask(taskID string, tenantID string) error {
	result := s.DB.Where("id = ? AND tenant_id = ?", taskID, tenantID).Delete(&models.Task{})
	if result.RowsAffected == 0 {
		return gorm.ErrRecordNotFound
	}
	return result.Error
}