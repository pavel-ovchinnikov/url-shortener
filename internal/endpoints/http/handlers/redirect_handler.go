package handlers

import (
	"net/http"
)

type RedirectHandler struct{}

func (h *RedirectHandler) Handler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNotImplemented)
	w.Write([]byte("RedirectHandler is not implemented yet"))
}
