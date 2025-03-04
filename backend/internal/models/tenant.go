package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Tenant struct {
	ID        uuid.UUID `gorm:"type:uuid;primaryKey;" json:"id"`
	Name      string    `gorm:"unique;not null" json:"name"`
	Users     []User    `gorm:"foreignKey:TenantID" json:"users,omitempty"`
	Tasks     []Task    `gorm:"foreignKey:TenantID" json:"tasks,omitempty"`
	CreatedAt time.Time `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt time.Time `gorm:"autoUpdateTime" json:"updated_at"`
}

func (t *Tenant) BeforeCreate(tx *gorm.DB) (err error) {
	t.ID = uuid.New()
	return nil
}
