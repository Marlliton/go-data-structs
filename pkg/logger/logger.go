package logger

import (
	"log/slog"
	"os"

	"github.com/Marlliton/slogpretty"
)

func Init() {
	opts := slogpretty.Options{
		Level:     slog.LevelDebug,
		Colorful:  true,
		Multiline: true,
	}
	handler := slogpretty.New(os.Stdout, &opts)
	logger := slog.New(handler)
	slog.SetDefault(logger)
}
