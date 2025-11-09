package handlers

import (
	"net/http"
)

type DeleteShortUrlHandler struct{}

func (h *DeleteShortUrlHandler) Handler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNotImplemented)
	w.Write([]byte("DeleteShortUrlHandler is not implemented yet"))
}
