package model

import (
	"time"
)

type UserAvatar struct {
	UserID    UUID   `gorm:"primary_key"`
	Bucket    string `gorm:"primary_key"`
	CreatedAt time.Time
	UpdatedAt time.Time
}
