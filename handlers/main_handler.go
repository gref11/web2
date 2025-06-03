package handlers

import (
	"net/http"

	"web3/internal/auth"
	"web3/internal/storage"
	"web3/scripts"
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
	scripts.EnableCORS(w)
	path := r.URL.Path

	switch path {
	case "/signup":
		h.RegistrationHandler(w, r)
	case "/success_signup":
		h.AfterRegistrationPageHandler(w, r)
	case "/login":
		h.LoginPageHandler(w, r)
	case "/User/add/":
		h.UserAddHandler(w, r)
	case "/User/login/":
		h.LoginHandler(w, r)
	case "/":
		auth.AuthMiddleware(h.UserPageHandler)(w, r)
	case "/User/update/":
		auth.AuthMiddleware(h.UserUpdateHandler)(w, r)
	default:
		h.StaticHandler(w, r)
	}
}
