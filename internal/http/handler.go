package http

import (
	"net/http"

	"github.com/isutare412/goasis/pkg/oapi"
)

type handler struct{}

func (h *handler) ListCafes(w http.ResponseWriter, r *http.Request) {
	// TODO: implement this
	w.WriteHeader(http.StatusNotImplemented)
}

func (h *handler) CreateCafe(w http.ResponseWriter, r *http.Request) {
	// TODO: implement this
	w.WriteHeader(http.StatusNotImplemented)
}

func (h *handler) DeleteCafe(w http.ResponseWriter, r *http.Request, cafeId oapi.PathCafeId) {
	// TODO: implement this
	w.WriteHeader(http.StatusNotImplemented)
}

func (h *handler) GetCafe(w http.ResponseWriter, r *http.Request, cafeId oapi.PathCafeId) {
	// TODO: implement this
	w.WriteHeader(http.StatusNotImplemented)
}

func (h *handler) ReplaceCafe(w http.ResponseWriter, r *http.Request, cafeId oapi.PathCafeId) {
	// TODO: implement this
	w.WriteHeader(http.StatusNotImplemented)
}
