package db

import (
	"context"
	"fmt"
	"log/slog"
	"os"
	"os/exec"

	"mickamy.com/playground/config"
)

func Migrate(ctx context.Context) error {
	cfg, err := config.DB().ParsedDBConfig()
	if err != nil {
		return fmt.Errorf("failed to parse database config: %w", err)
	}

	// mysqldef -u $DB_USER -p$DB_PASSWORD -h $DB_HOST -P $DB_PORT $DB_NAME
	cmd := exec.CommandContext(ctx,
		"mysqldef",
		"-u", cfg.User,
		"-p"+cfg.Pass,
		"-h", cfg.Host,
		"-P", cfg.Port,
		cfg.Name,
	)

	schemaFile, err := os.Open(config.Common().PackageRoot + "/db/schema.sql")
	if err != nil {
		return fmt.Errorf("failed to open schema file: %w", err)
	}
	defer func(f *os.File) {
		err := f.Close()
		if err != nil {
			slog.Warn("failed to close schema file", "err", err)
		}
	}(schemaFile)
	cmd.Stdin = schemaFile

	output, err := cmd.CombinedOutput()
	if err != nil {
		return fmt.Errorf("failed to run migration: %w\nOutput: %s", err, string(output))
	}

	return nil
}
