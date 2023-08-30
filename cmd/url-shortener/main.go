package main

import (
	"fmt"

	"url-shortener/internal/config"
)

func main() {
	cfg := config.MustLoad()
	fmt.Println(cfg)

	// TODO: init logger: slong
	// TODO: init storage
	// TODO: init router: chi
	// TODO: init server
}
