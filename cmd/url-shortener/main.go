package main

import (
	"fmt"
	"io"
	"net/http"

	"url-shortener/internal/config"
)

func hello(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, "This is my website!\n")
	fmt.Println("Hello")
}

func main() {
	cnf := config.Config{}
	cnf.Address = ":8090"

	http.HandleFunc("/hello", hello)
	http.ListenAndServe(cnf.HTTPServer.Address, nil)
}
