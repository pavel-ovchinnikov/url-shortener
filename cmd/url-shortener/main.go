package main

import (
	"log"
	"net/http"
	"os"

	"url-shortener/internal/config"
	"url-shortener/internal/http-server/handler/redirect"
	"url-shortener/internal/http-server/handler/save"
	"url-shortener/internal/storage/sqlite"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func main() {
	// TODO: init logger: slong
	// TODO: add func storage.DeleteUrl
	// TODO: add template html
	// TODO: add autheration
	// TODO: handler
	//		DeleteURL
	//		GaveURL
	// TODO: test handler
	//		SaveURL
	//		GaveURL
	//		DeleteURL
	cfg := config.MustLoad()

	storage, err := sqlite.New(cfg.StoragePath)
	if err != nil {
		log.Fatalf("failed to init storage %s", err)
		os.Exit(1)
	}

	router := chi.NewRouter()
	router.Use(middleware.RequestID)
	router.Use(middleware.RequestID)
	router.Use(middleware.Recoverer)
	router.Use(middleware.URLFormat)
	router.Post("/url", save.New(storage))
	router.Get("/{alias}", redirect.New(storage))
	// router.Get("/url/{alias}", redirect.New(storage))

	log.Println("start server", "address", cfg.Address)

	srv := &http.Server{
		Addr:         cfg.Address,
		Handler:      router,
		ReadTimeout:  cfg.HTTPServer.Timeout,
		WriteTimeout: cfg.HTTPServer.Timeout,
		IdleTimeout:  cfg.HTTPServer.IdleTimeout,
	}

	if err := srv.ListenAndServe(); err != nil {
		log.Println("failed to start server")
	}
}
