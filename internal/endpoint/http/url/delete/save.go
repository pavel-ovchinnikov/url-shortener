package delete

import (
	"net/http"
)

type URLDelete struct {
}

func (h *URLDelete) Handler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusInternalServerError)
	w.Write([]byte("500 - Something bad happened!"))
}

func NewURLDelete() *URLDelete {
	return &URLDelete{}
}
