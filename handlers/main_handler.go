package handlers

import (
	"net/http"

	"web3/internal/storage"
)

type Handler struct {
	Storage storage.Storage
}

func NewHandler() (*Handler, error) {
	st, err := storage.NewDataBaseStorage()
	if err != nil {
		return &Handler{}, err
	}
	// var st storage.Storage

	h := &Handler{
		Storage: st,
	}

	return h, nil
}

func (h *Handler) MainHandler(w http.ResponseWriter, r *http.Request) {
	path := r.URL.Path

	switch path {
	case "/":
		h.RegistrationHandler(w, r)
	case "/User/add/":
		h.UserAddHandler(w, r)
	default:
		h.StaticHandler(w, r)
	}
}
