package save

import (
	"errors"
	"io"
	"log"
	"net/http"

	resp "url-shortener/internal/lib/api/response"
	"url-shortener/internal/storage"

	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/render"
	"github.com/go-playground/validator/v10"
)

type Request struct {
	URL   string `json:"url" validate:"requred,url"`
	Alias string `json:"alias,omitempty"`
}

type Response struct {
	resp.Response
	Alias string `json:"alias,omitempty"`
}

type URLSave interface {
	SaveURL(urlToSave string, alias string) (int64, error)
}

func New(urlSaver URLSave) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		const op = "handler.url.Save.New"

		log.Println(op, "reques_id", middleware.GetReqID(r.Context()))

		var req Request

		err := render.DecodeJSON(r.Body, &req)
		if errors.Is(err, io.EOF) {
			log.Println("request body is empty")
			render.JSON(w, r, resp.Error("empty request"))
			return
		}

		if err != nil {
			log.Println("failed to decode request body", err)
			render.JSON(w, r, resp.Error("failed to decode request"))
			return
		}

		log.Println("request body decoded", "request", req)

		if err := validator.New().Struct(req); err != nil {
			validateErr := err.(validator.ValidationErrors)
			log.Println("invalid request", err)
			render.JSON(w, r, resp.ValidationError(validateErr))
			return
		}

		alias := req.Alias
		// if alias == "" {
		// 	alias = random.NewRandomString(aliasLength)
		// }

		id, err := urlSaver.SaveURL(req.URL, alias)
		if errors.Is(err, storage.ErrURLExists) {
			log.Println("url already exists", "url", req.URL)
			render.JSON(w, r, resp.Error("url already exists"))
			return
		}

		if err != nil {
			log.Println("failed to add url", err)
			render.JSON(w, r, resp.Error("failed to add url"))
			return
		}

		log.Println("url added", "id", id)

		// response OK
		render.JSON(w, r, Response{
			Response: resp.OK(),
			Alias:    alias,
		})
	}
}
