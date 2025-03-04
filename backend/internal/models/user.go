package models

import (
	"fmt"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

// Role constants for RBAC
const (
	RoleAdmin   = "admin"
	RoleManager = "manager"
	RoleUser    = "user"
)

type User struct {
	ID        uuid.UUID `gorm:"type:uuid;primaryKey;" json:"id"`
	Name      string    `gorm:"not null" json:"name"`
	Email     string    `gorm:"unique;not null" json:"email"`
	Password  string    `gorm:"not null" json:"-"`
	TenantID  uuid.UUID `gorm:"type:uuid;not null;index" json:"tenant_id"`
	Role      string    `gorm:"type:varchar(10);not null" json:"role"`
	Tenant    Tenant    `gorm:"foreignKey:TenantID" json:"tenant,omitempty"`
	CreatedAt time.Time `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt time.Time `gorm:"autoUpdateTime" json:"updated_at"`
}

// BeforeCreate hook to auto-generate UUIDs
func (u *User) BeforeCreate(tx *gorm.DB) (err error) {
	u.ID = uuid.New()
	return nil
}

// CreateUser inserts a new user into the database
func CreateUser(user *User, db *gorm.DB) error {
	user.ID = uuid.New()
	result := db.Create(user)
	return result.Error
}

// FindUserByID retrieves a user by their ID
func FindUserByID(id uuid.UUID, db *gorm.DB) (*User, error) {
	var user User
	result := db.Preload("Tenant").Where("id = ?", id).First(&user)
	if result.Error != nil {
		if result.Error == gorm.ErrRecordNotFound {
			return nil, nil
		}
		return nil, result.Error
	}
	fmt.Printf("User found: %+v\n", user)
	return &user, nil
}

// FindUserByEmail retrieves a user by their email
func FindUserByEmail(email string, db *gorm.DB) (*User, error) {
	var user User
	result := db.Preload("Tenant").Where("email = ?", email).First(&user)
	if result.Error != nil {
		if result.Error == gorm.ErrRecordNotFound {
			return nil, nil
		}
		fmt.Printf("Unexpected error: %v\n", result.Error)
		return nil, result.Error
	}
	fmt.Printf("User found: %+v\n", user)
	return &user, nil
}
