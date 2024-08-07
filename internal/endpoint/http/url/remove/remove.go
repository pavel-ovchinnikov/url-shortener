package remove

import (
	"context"
	"log/slog"
	"net/http"
)

type URLStorage interface {
	DeleteURL(ctx context.Context, url string) error
}

type URLRemove struct {
	storage URLStorage
	log     *slog.Logger
}

func (h *URLRemove) Handler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusInternalServerError)
	w.Write([]byte("500 - Something bad happened!"))
}

func NewURLRemove(log *slog.Logger, storage URLStorage) *URLRemove {
	return &URLRemove{
		storage: storage,
		log:     log,
	}
}
