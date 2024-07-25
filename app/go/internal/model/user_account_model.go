package model

import (
	"time"
)

type UserAccount struct {
	UserID    UUID `gorm:"primaryKey"`
	User      User
	Email     string
	Provider  UserAccountProvider `gorm:"type:enum('google')"`
	UID       string
	CreatedAt time.Time
	UpdatedAt time.Time
}
