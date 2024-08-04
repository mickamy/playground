package test

import (
	"context"
	"fmt"
	"log/slog"
	"path/filepath"
	"testing"
	"time"

	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/api/types/mount"
	"github.com/google/uuid"
	"github.com/testcontainers/testcontainers-go"
	"github.com/testcontainers/testcontainers-go/wait"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	gormLogger "gorm.io/gorm/logger"

	"mickamy.com/playground/config"
	"mickamy.com/playground/internal/lib/logger"
)

const (
	useTestContainers = true
	reuseContainer    = false
)

func NewTestDB(t *testing.T) *gorm.DB {
	t.Helper()

	if useTestContainers {
		return initTestContainers(t)
	}

	return initActualDB(t)
}

func initTestContainers(t *testing.T) *gorm.DB {
	ctx := context.Background()

	packageRoot := config.Common().PackageRoot
	ctn, err := testcontainers.GenericContainer(ctx, testcontainers.GenericContainerRequest{
		ContainerRequest: testcontainers.ContainerRequest{
			Name:         uuid.NewString(),
			Image:        "mysql:8.0.36",
			ExposedPorts: []string{"3306/tcp"},
			Env: map[string]string{
				"MYSQL_ROOT_PASSWORD": "root",
				"MYSQL_USER":          "user",
				"MYSQL_PASSWORD":      "pass",
				"MYSQL_DATABASE":      "playground",
			},
			HostConfigModifier: func(hostConfig *container.HostConfig) {
				hostConfig.Mounts = append(hostConfig.Mounts, mount.Mount{
					Type:        mount.TypeBind,
					Source:      filepath.Join(packageRoot, "db", "schema.sql"),
					Target:      "/docker-entrypoint-initdb.d/init.sql",
					ReadOnly:    true,
					BindOptions: nil,
				})
			},
			WaitingFor: wait.ForLog("port: 3306  MySQL Community Server - GPL").
				WithStartupTimeout(15 * time.Second),
		},
		Started: true,
		Reuse:   reuseContainer,
		Logger:  logger.NewLogger(slog.LevelInfo),
	})
	if err != nil {
		t.Fatalf("Could not start mysql: %s", err)
	}

	host, err := ctn.Host(ctx)
	if err != nil {
		t.Fatalf("Could not get host: %s", err)
	}
	port, err := ctn.MappedPort(ctx, "3306")
	if err != nil {
		t.Fatalf("Could not get port %s: %s", "3306", err)
	}

	t.Cleanup(func() {
		if err := ctn.Terminate(ctx); err != nil {
			t.Fatalf("Could not stop mysql: %s", err)
		}
	})

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?parseTime=true&loc=Local", "user", "pass", host, port.Int(), "playground")
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: gormLogger.Default.LogMode(gormLogger.Info),
	})
	if err != nil {
		t.Fatalf("Could not connect to database: %s", err)
	}

	return db
}

func initActualDB(t *testing.T) *gorm.DB {
	db, err := gorm.Open(mysql.Open(config.DB().URL), &gorm.Config{
		Logger: gormLogger.Default.LogMode(gormLogger.Warn),
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
