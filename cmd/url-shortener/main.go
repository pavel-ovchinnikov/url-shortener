package main

import (
	"log"
	"os"

	"url-shortener/internal/config"
	"url-shortener/internal/storage/sqlite"
)

func main() {

	cfg := config.MustLoad()

	storage, err := sqlite.New(cfg.StoragePath)
	if err != nil {
		log.Fatalf("failed to init storage %s", err)
		os.Exit(1)
	}

	id, err := storage.SaveURL("https:/google.com", "google")
	if err != nil {
		log.Fatalf("failed to init storage %s", err)
	} else {
		log.Println("id ", id)
	}

	// TODO: init logger: slong
	// TODO: init router: chi
	// TODO: init server
}
