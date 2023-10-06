package handler

import (
	"github.com/go-chi/chi/v5"
	serviceCache "github.com/patrickmn/go-cache"
	"github.com/ziggsdil/zero-level-wb/pkg/db"
	"github.com/ziggsdil/zero-level-wb/pkg/renderer"
)

type Handler struct {
	db *db.Database
	c  *serviceCache.Cache

	renderer renderer.Renderer
}

func NewHandler(db *db.Database, c *serviceCache.Cache) *Handler {
	return &Handler{
		db: db,
		c:  c,
	}
}

func (h *Handler) Router() chi.Router {
	router := chi.NewRouter()

	router.Route("/api", func(r chi.Router) {
		r.Route("/info", func(r chi.Router) {
			r.Get("/{orderID}", h.GetInfo)
		})
	})

	return router
}
