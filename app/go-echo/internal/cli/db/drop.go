package db

import (
	"context"
	"database/sql"
	"fmt"
	"log/slog"

	_ "github.com/go-sql-driver/mysql"

	"mickamy.com/playground/config"
)

func Drop(ctx context.Context) error {
	dsn := config.DB().URL
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
	if _, err := db.ExecContext(ctx, `DROP DATABASE IF EXISTS playground`); err != nil {
		return fmt.Errorf("failed to drop database: %w", err)
	}

	return nil
}
