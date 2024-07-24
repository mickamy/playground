package database

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"

	"mickamy.com/playground/config"
)

var db *gorm.DB

func DB(cfg config.DBConfig) *gorm.DB {
	if db == nil {
		initializeDB(cfg)
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
