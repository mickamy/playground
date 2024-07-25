package model

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

var (
	UserSlugReserved = []string{"me", "example", "terms", "tos", "privacy"}
)

type User struct {
	ID        UUID
	Slug      string
	CreatedAt time.Time
	UpdatedAt time.Time
}

func (u *User) BeforeCreate(tx *gorm.DB) error {
	id, err := uuid.NewRandom()
	u.ID = UUID(id)
	return err
}
