package handler

import (
	"fmt"
	"github.com/go-chi/chi/v5"
	"net/http"
)

func (h *Handler) GetInfo(w http.ResponseWriter, r *http.Request) {
	// todo: implement
	orderID := chi.URLParam(r, "orderID")
	if orderID == "" {
		h.renderer.RenderError(w, fmt.Errorf("order id param is empty\n"))
		return
	}

	data, ok := h.c.Get(orderID)
	if !ok {
		h.renderer.RenderError(w, fmt.Errorf("failed to find key: %s", orderID))
		return
	}

	h.renderer.RenderOK(w)

	h.renderer.RenderJSON(w, data)
}
