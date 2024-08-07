package main

import (
	"context"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/go-chi/chi/v5"

	"github.com/pavel-ovchinnikov/url-shortener/internal/config"
	handler_redirect "github.com/pavel-ovchinnikov/url-shortener/internal/endpoint/http/redirect"
	handler_remove "github.com/pavel-ovchinnikov/url-shortener/internal/endpoint/http/url/remove"
	handler_save "github.com/pavel-ovchinnikov/url-shortener/internal/endpoint/http/url/save"
	url_storage "github.com/pavel-ovchinnikov/url-shortener/internal/storage/url_storage"
)

func main() {
	// init config
	cfg := config.MustLoad()

	// init context
	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer cancel()

	// init log
	log := NewLogger()

	// init storage
	dbStorage, err := url_storage.NewURLStorage(&cfg.URLStorage)
	if err != nil {
		log.Error(err.Error())
		return
	}
	_ = dbStorage

	// init server
	httpServer := NewHTTPServer(ctx, log, &cfg.HTTPServer, dbStorage)

	// run server
	if err := httpServer.ListenAndServe(); err != nil {
		log.Error(err.Error())
	}
}

func NewLogger() *slog.Logger {
	return slog.New(
		slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug}),
	)
}

func NewHTTPServer(
	ctx context.Context,
	log *slog.Logger,
	cfg *config.HTTPServer,
	dbStorage *url_storage.URLStorage,
) *http.Server {
	router := chi.NewRouter()
	router.Route("/url", func(r chi.Router) {
		r.Post("/", handler_save.NewURLSave(log, dbStorage).Handler)
		r.Delete("/", handler_remove.NewURLRemove(log, dbStorage).Handler)
	})
	router.Get("/{alias}", handler_redirect.NewRedirect(log, dbStorage).Handler)

	return &http.Server{
		Addr:         cfg.Address,
		Handler:      router,
		ReadTimeout:  cfg.Timeout,
		WriteTimeout: cfg.Timeout,
		IdleTimeout:  cfg.Timeout,
	}
}
