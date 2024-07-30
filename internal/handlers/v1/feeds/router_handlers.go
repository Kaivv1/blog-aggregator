package feeds

import (
	"github.com/Kaivv1/blog-aggregator/internal/config"
	"github.com/Kaivv1/blog-aggregator/internal/handlers/v1/auth"
	"github.com/go-chi/chi"
)

type feedsHandler struct {
	config *config.ApiConfig
}

func newFeedsHandler(cfg *config.ApiConfig) *feedsHandler {
	return &feedsHandler{
		config: cfg,
	}
}

func NewFeedsRouter(cfg *config.ApiConfig) *chi.Mux {
	m := auth.NewMiddleware(cfg)
	handler := newFeedsHandler(cfg)
	router := chi.NewRouter()
	router.Post("/", m.AuthMiddleware(handler.CreateFeed))
	router.Get("/", handler.GetFeeds)

	return router
}
