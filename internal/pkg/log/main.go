package log

import (
	"log/slog"
	"os"

	"github.com/zeon-code/tiny-url/internal/pkg/config"
)

func NewLogger(c config.Configuration) *slog.Logger {
	var level slog.Level

	switch string(c.Log().Level()) {
	case "debug":
		level = slog.LevelDebug
	case "info":
		level = slog.LevelInfo
	case "warn":
		level = slog.LevelWarn
	case "error":
		level = slog.LevelError
	default:
		level = slog.LevelInfo
	}

	return slog.New(slog.NewJSONHandler(
		os.Stdout,
		&slog.HandlerOptions{
			Level: level,
		},
	))
}
