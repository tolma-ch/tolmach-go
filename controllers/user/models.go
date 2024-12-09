package user

import (
	"time"
)

type User struct {
	ID          uint   // Standard field for the primary key
	Username    string // A regular string field
	FirstName   string
	LastName    string
	Password    string
	Email       *string // A pointer to a string, allowing for null values
	IsActive    bool
	IsStaff     bool
	IsSuperuser bool
	LastLogin   time.Time
	CreatedAt   time.Time // Automatically managed by GORM for creation time
	UpdatedAt   time.Time // Automatically managed by GORM for update time
}
