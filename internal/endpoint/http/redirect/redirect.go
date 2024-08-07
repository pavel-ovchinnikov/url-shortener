package redirect

import (
	"context"
	"log/slog"
	"net/http"

	"github.com/go-chi/chi/v5"
)

type URLStorage interface {
	GetURL(ctx context.Context, alias string) (string, error)
}

type Redirect struct {
	storage URLStorage
	log     *slog.Logger
}

func (h *Redirect) Handler(w http.ResponseWriter, r *http.Request) {
	resURL := "https://ru.wikipedia.org/wiki/IBM"
	alias := chi.URLParam(r, "alarm")
	if alias == "" {
		h.log.Error("Empty alias")
		w.WriteHeader(http.StatusNotFound)
		return
	}

	resURL, err := h.storage.GetURL(r.Context(), alias)
	if err != nil {
		h.log.Error(err.Error())
		return
	}

	http.Redirect(w, r, resURL, http.StatusFound)
}

func NewRedirect(log *slog.Logger, storage URLStorage) *Redirect {
	return &Redirect{
		storage: storage,
		log:     log,
	}
}
