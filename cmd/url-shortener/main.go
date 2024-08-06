package main

import (
	"context"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/pavel-ovchinnikov/url-shortener/internal/config"
	"github.com/pavel-ovchinnikov/url-shortener/internal/storage/urlstorage"
)

func main() {
	// init config
	cfg := config.MustLoad()

	// init context
	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer cancel()

	// init logger
	log := NewLogger()

	// init storage
	dbStorage, err := urlstorage.NewURLStorage(&cfg.URLStorage)
	if err != nil {
		log.Error(err.Error())
		return
	}
	_ = dbStorage

	// init server
	httpServer := NewHTTPServer(ctx, &cfg.HTTPServer)

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

func NewHTTPServer(ctx context.Context, cfg *config.HTTPServer) *http.Server {
	handlerFunc := func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("500 - Something bad happened!"))
	}

	router := chi.NewRouter()
	router.Route("/url", func(r chi.Router) {
		r.Post("/", handlerFunc)
		r.Delete("/", handlerFunc)
	})
	router.Get("/{alias}", handlerFunc)

	return &http.Server{
		Addr:         cfg.Address,
		Handler:      router,
		ReadTimeout:  time.Duration(time.Second * 2),
		WriteTimeout: time.Duration(time.Second * 2),
		IdleTimeout:  time.Duration(time.Second * 2),
	}
}
