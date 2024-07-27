package database

import (
	"sync"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"

	"mickamy.com/playground/config"
)

var (
	once sync.Once
	db   *gorm.DB
)

func DB(cfg config.DBConfig) *gorm.DB {
	if db == nil {
		once.Do(func() {
			initializeDB(cfg)
		})
	}
	return db
}

func initializeDB(cfg config.DBConfig) {
	var err error
	db, err = gorm.Open(mysql.Open(cfg.DatabaseURL()), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Warn),
	})
	if err != nil {
		panic(err)
	}
}
