package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

const (
	StatusPending   = "pending"
	StatusCompleted = "completed"
)

// Task represents a task in the system
type Task struct {
	ID          uuid.UUID `gorm:"type:uuid;primaryKey;" json:"id"`
	Title       string    `gorm:"not null" json:"title"`
	Description string    `json:"description"`
	Status      string    `gorm:"type:varchar(20);default:'pending'" json:"status"`
	TenantID    uuid.UUID `gorm:"type:uuid;not null;index" json:"tenant_id"`
	Tenant      Tenant    `gorm:"foreignKey:TenantID" json:"tenant,omitempty"`
	CreatedAt   time.Time `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt   time.Time `gorm:"autoUpdateTime" json:"updated_at"`
}

// BeforeCreate hook to set UUID before inserting into the database
func (t *Task) BeforeCreate(tx *gorm.DB) (err error) {
	t.ID = uuid.New()
	if t.Status == "" {
		t.Status = StatusPending
	}
	return nil
}
