package base_model

import (
	"time"
)

type AbstractBase struct {
	ID        string    `gorm:"type:uuid;primaryKey;unique;not null" json:"id"`
	CreatedAt time.Time `gorm:"column:created_at; not null; autoCreateTime" json:"created_at"`
	UpdatedAt time.Time `gorm:"column:updated_at; null; autoUpdateTime" json:"updated_at"`
}
