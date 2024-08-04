package db

import (
	"context"
	"database/sql"
	"fmt"
	"log/slog"

	_ "github.com/go-sql-driver/mysql"

	"mickamy.com/playground/config"
)

func Create(ctx context.Context) error {
	cfg, err := config.DB().ParsedDBConfig()
	if err != nil {
		return fmt.Errorf("failed to parse DB config: %w", err)
	}
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/?parseTime=true&loc=Local", cfg.User, cfg.Pass, cfg.Host, cfg.Port)
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		return fmt.Errorf("failed to open database connection: %w", err)
	}

	defer func(db *sql.DB) {
		err := db.Close()
		if err != nil {
			slog.Warn("failed to close DB connection", "err", err)
		}
	}(db)

	if _, err := db.ExecContext(ctx, `CREATE DATABASE IF NOT EXISTS playground`); err != nil {
		return fmt.Errorf("failed to create database: %w", err)
	}

	return nil
}
