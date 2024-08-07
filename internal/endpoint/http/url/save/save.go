package save

import (
	"context"
	"log/slog"
	"net/http"
)

type URLStorage interface {
	SaveURL(ctx context.Context, url string, alias string) (int64, error)
}

type URLSave struct {
	storage URLStorage
	log     *slog.Logger
}

func (h *URLSave) Handler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusInternalServerError)
	w.Write([]byte("500 - Something bad happened!"))
}

func NewURLSave(log *slog.Logger, storage URLStorage) *URLSave {
	return &URLSave{
		storage: storage,
		log:     log,
	}
}
