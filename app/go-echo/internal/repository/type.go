package repository

import (
	"context"

	"gorm.io/gorm"
)

type Scope = func(*gorm.DB) *gorm.DB

type Creator[MODEL any] interface {
	Create(ctx context.Context, m *MODEL) error
}

type Reader[MODEL any, ID any] interface {
	Get(ctx context.Context, id ID, scopes ...Scope) (MODEL, error)
}

type Updater[MODEL any] interface {
	Update(ctx context.Context, m MODEL) error
}

type Deleter[ID any] interface {
	Delete(ctx context.Context, id ID) error
}
