package sl

import (
	"log/slog"
	"os"

	"github.com/lmittmann/tint"
)


const (
	envLocal = "local"
	envDev = "dev"
	envProd = "prod"
)

func SetUpLogger(env string) *slog.Logger {
	var level slog.Level

	switch env {
	case envLocal:
		level = slog.LevelDebug
	case envDev:
		level = slog.LevelInfo
	case envProd:
		level = slog.LevelWarn
	default:
		level = slog.LevelInfo 
	}

	handler := tint.NewHandler(os.Stdout, &tint.Options{Level: level})
	return slog.New(handler)
}