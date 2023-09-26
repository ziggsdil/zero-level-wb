package handler

import (
	"github.com/go-chi/chi/v5"
	"github.com/ziggsdil/zero-level-wb/pkg/db"
)

type Handler struct {
	db *db.Database
}

func NewHandler(db *db.Database) *Handler {
	return &Handler{
		db: db,
	}
}

func (h *Handler) Router() chi.Router {
	router := chi.NewRouter()

	router.Route("/", func(r chi.Router) {
		r.Get("/add", h.GetInfo)
	})

	return router
}
