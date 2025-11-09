package main

import (
	"fmt"

	"github.com/pavel-ovchinnikov/url-shortener/internal/config"
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

	// TODO: init http server

	// TODO: start http server
}
