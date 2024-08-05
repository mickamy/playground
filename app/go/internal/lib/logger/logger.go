package logger

import (
	"log"
	"log/slog"
	"os"
)

func init() {
	opts := &slog.HandlerOptions{
		AddSource: true,
		Level:     slog.LevelDebug,
	}
	handler := slog.NewTextHandler(os.Stdout, opts)
	slog.SetDefault(slog.New(handler))
}

func NewLogger(level slog.Level) *log.Logger {
	opts := &slog.HandlerOptions{
		AddSource: true,
		Level:     level,
	}
	handler := slog.NewTextHandler(os.Stdout, opts)
	return slog.NewLogLogger(handler, level)
}
