package utils

import (
	"log/slog"
	"os"
)

func InitLogger() *slog.Logger {
	logger := slog.New(slog.NewTextHandler(os.Stdout, nil))
	return logger
}
