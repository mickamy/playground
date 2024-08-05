package di

import (
	"github.com/google/wire"
	"gorm.io/gorm"

	"mickamy.com/playground/internal/infra/store/database"
)

type Infras struct {
	DB *gorm.DB
}

var infras = wire.NewSet(
	database.DB,
)
