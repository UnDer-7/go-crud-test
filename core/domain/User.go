package domain

import (
	"gorm.io/gorm"
	"time"
)

type User struct {
	gorm.Model
	ID       int
	Email    string
	Password string
	CreatedAt    time.Time
	UpdatedAt    time.Time
}
