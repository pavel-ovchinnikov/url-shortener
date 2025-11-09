package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/pavel-ovchinnikov/url-shortener/internal/config"
	"github.com/pavel-ovchinnikov/url-shortener/internal/endpoints/http/handlers"
)

func main() {
	cfg, err := config.LoadConfig()
	if err != nil {
		panic("Failed to load config: " + err.Error())
	}

	fmt.Println("Config loaded:", cfg)

	// TODO: init logger

	// TODO: init database

	// TODO: init cache

	mux := http.NewServeMux()
	handlers.RegisterHandlers(mux)

	fmt.Printf("Server started at %s\n", cfg.HTTPServer.Address)
	log.Fatal(http.ListenAndServe(cfg.HTTPServer.Address, mux))
}
