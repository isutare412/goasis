package http

import (
	"net/http"

	oapitypes "github.com/oapi-codegen/runtime/types"
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

func (h *handler) DeleteCafe(w http.ResponseWriter, r *http.Request, cafeId oapitypes.UUID) {
	// TODO: implement this
	w.WriteHeader(http.StatusNotImplemented)
}

func (h *handler) GetCafe(w http.ResponseWriter, r *http.Request, cafeId oapitypes.UUID) {
	// TODO: implement this
	w.WriteHeader(http.StatusNotImplemented)
}

func (h *handler) ReplaceCafe(w http.ResponseWriter, r *http.Request, cafeId oapitypes.UUID) {
	// TODO: implement this
	w.WriteHeader(http.StatusNotImplemented)
}
