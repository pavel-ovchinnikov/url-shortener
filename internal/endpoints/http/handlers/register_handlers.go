package handlers

import (
	"net/http"
)

func RegisterHandlers(mux *http.ServeMux) {
	createHandler := &CreateShortUrlHandler{}
	deleteHandler := DeleteShortUrlHandler{}
	redirectHandler := &RedirectHandler{}

	mux.HandleFunc("/create", createHandler.Handler)
	mux.HandleFunc("/delete", deleteHandler.Handler)
	mux.HandleFunc("/redirect", redirectHandler.Handler)
}
