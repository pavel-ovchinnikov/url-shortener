package main

import (
	"fmt"
	"io"
	"net/http"

	"url-shortener/internal/config"
)

func addLink(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, "This is my website!\n")
	fmt.Println("Hello")
}

func removeLink(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, "This is my website!\n")
	fmt.Println("Hello")
}

func patchLink(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, "This is my website!\n")
	fmt.Println("Hello")
}

func main() {
	cnf := config.Config{}
	cnf.Address = ":8090"

	http.HandleFunc("/add", addLink)
	http.HandleFunc("/remove", removeLink)
	http.HandleFunc("/patch", patchLink)
	http.ListenAndServe(cnf.HTTPServer.Address, nil)
}
