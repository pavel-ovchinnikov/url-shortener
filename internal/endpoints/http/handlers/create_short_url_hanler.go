package handlers

import (
	"net/http"
)

type CreateShortUrlHandler struct{}

func (h *CreateShortUrlHandler) Handler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNotImplemented)
	w.Write([]byte("CreateShortUrlHandler is not implemented yet"))
}
