package feedfollow

import (
	"github.com/Kaivv1/blog-aggregator/internal/config"
	"github.com/Kaivv1/blog-aggregator/internal/handlers/v1/auth"
	"github.com/go-chi/chi"
)

type feedFollowHandler struct {
	config *config.ApiConfig
}

func newFeedFollowHandler(cfg *config.ApiConfig) *feedFollowHandler {
	return &feedFollowHandler{
		config: cfg,
	}
}
func NewFeedFollowRouter(cfg *config.ApiConfig) *chi.Mux {
	handler := newFeedFollowHandler(cfg)
	router := chi.NewRouter()
	m := auth.NewMiddleware(cfg)
	router.Post("/", m.AuthMiddleware(handler.CreateFeedFolow))
	router.Delete("/{ff_id}", m.AuthMiddleware(handler.DeleteFeedFollow))
	router.Get("/", m.AuthMiddleware(handler.GetFollowedFeedsUser))
	return router
}
