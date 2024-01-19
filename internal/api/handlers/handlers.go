package handlers

import "net/http"

type IHandlers interface {
	Hi(w http.ResponseWriter, r *http.Request)
}

func NewHandler() IHandlers {
	return &handler{todo: false}
}

type handler struct {
	todo bool
}

func (h *handler) Hi(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("You're welcome"))
}
