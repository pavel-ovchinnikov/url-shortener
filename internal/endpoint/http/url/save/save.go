package save

import (
	"net/http"
)

type URLSave struct {
}

func (h *URLSave) Handler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusInternalServerError)
	w.Write([]byte("500 - Something bad happened!"))
}

func NewURLSave() *URLSave {
	return &URLSave{}
}
