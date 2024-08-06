package main

import (
	"log/slog"
	"os"

	"github.com/pavel-ovchinnikov/url-shortener/internal/config"
)

func main() {
	// TODO: init config
	cfg := config.MustLoad()
	_ = cfg

	// TODO: init context
	// TODO: init logger
	log := NewLogger()

	// TODO: init storage
	// TODO: init router
	// TODO: run server
}

func NewLogger() *slog.Logger {
	return slog.New(
		slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug}),
	)
}
