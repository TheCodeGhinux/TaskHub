

package models

import "time"

type UserProfileResponse struct {
    ProfileID  string    `json:"profile_id"`
    FirstName  string    `json:"first_name"`
    LastName   string    `json:"last_name"`
    FullName   string    `json:"full_name"`
    UserName   string    `json:"user_name"`
    Phone      string    `json:"phone"`
    AvatarURL  string    `json:"avatar_url"`
    UserID     string    `json:"user_id"`
    CreatedAt  time.Time `json:"created_at"`
    UpdatedAt  time.Time `json:"updated_at"`
}

type UserResponse struct {
    ID         string      `json:"id"`
    Name       string      `json:"name"`
    Email      string      `json:"email"`
    IsVerified bool        `json:"is_verified"`
    Deactivated bool       `json:"deactivated"`
    IsActive   bool        `json:"is_active"`
    Profile    UserProfileResponse `json:"profile"`
    CreatedAt  time.Time   `json:"created_at"`
    UpdatedAt  time.Time   `json:"updated_at"`
    Role       string      `json:"role"`
}

// ErrorResponse represents a standard error response structure.
type ErrorResponse struct {
    Status     string `json:"status"`       // Status of the response, e.g., "Error"
    StatusCode int    `json:"status_code"`  // HTTP status code, e.g., 500
    Message    string `json:"message"`      // Error message, e.g., "Failed to check user existence"
}

