package model

import (
	"time"
)

type UserProfile struct {
	UserID    UUID `gorm:"primary_key"`
	User      User
	Name      string
	Bio       string
	CreatedAt time.Time
	UpdatedAt time.Time
}
