package model

import (
	"time"
)

var (
	UserSlugReserved = []string{"me", "example", "terms", "tos", "privacy"}
)

type User struct {
	ID        string
	Slug      string
	CreatedAt time.Time
	UpdatedAt time.Time
}
