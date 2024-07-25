package test

import (
	"context"
	"fmt"
	"os"
	"path/filepath"
	"testing"
	"time"

	"github.com/testcontainers/testcontainers-go"
	mysqlContainers "github.com/testcontainers/testcontainers-go/modules/mysql"
	"github.com/testcontainers/testcontainers-go/wait"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"

	"mickamy.com/playground/config"
)

const (
	UseTestContainers = true
)

func NewTestDB(t *testing.T) *gorm.DB {
	t.Helper()

	if UseTestContainers {
		return initTestContainers(t)
	}

	return initTestDB(t)
}

func initTestContainers(t *testing.T) *gorm.DB {
	ctx := context.Background()

	packageRoot, ok := os.LookupEnv("PACKAGE_ROOT")
	if !ok {
		t.Fatal("PACKAGE_ROOT environment variable not set")
	}
	cfg := config.DB()
	if packageRoot == "" {
		panic("PACKAGE_ROOT environment variable not set")
	}
	container, err := mysqlContainers.Run(ctx,
		"mysql:8.0.36",
		mysqlContainers.WithScripts(filepath.Join(packageRoot, "db", "schema.sql")),
		mysqlContainers.WithDatabase(cfg.Name),
		mysqlContainers.WithUsername(cfg.User),
		mysqlContainers.WithPassword(cfg.Password),
		testcontainers.WithWaitStrategy(
			wait.ForLog("port: 3306  MySQL Community Server - GPL").
				WithStartupTimeout(20*time.Second)),
	)
	if err != nil {
		t.Fatalf("Could not start mysql: %s", err)
	}

	host, err := container.Host(ctx)
	if err != nil {
		t.Fatalf("Could not get host: %s", err)
	}
	port, err := container.MappedPort(ctx, "3306")
	if err != nil {
		t.Fatalf("Could not get port %s: %s", "3306", err)
	}

	t.Cleanup(func() {
		if err := container.Terminate(ctx); err != nil {
			t.Fatalf("Could not stop mysql: %s", err)
		}
	})

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?parseTime=true&loc=Local", cfg.User, cfg.Password, host, port.Int(), cfg.Name)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		t.Fatalf("Could not connect to database: %s", err)
	}

	return db
}

func initTestDB(t *testing.T) *gorm.DB {
	db, err := gorm.Open(mysql.Open(config.DB().DatabaseURL()), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Warn),
	})
	if err != nil {
		panic(err)
	}

	t.Cleanup(func() {
		sqlDB, err := db.DB()
		if err != nil {
			t.Fatalf("Could not get DB connection: %s", err)
		}
		err = sqlDB.Close()
		if err != nil {
			t.Fatalf("Could not close DB connection: %s", err)
		}
	})

	return db
}
