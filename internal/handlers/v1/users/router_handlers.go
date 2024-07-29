package users

import (
	"github.com/Kaivv1/blog-aggregator/internal/config"
	"github.com/go-chi/chi"
)

type usersHandler struct {
	config *config.ApiConfig
}

func newUsersHandler(cfg *config.ApiConfig) *usersHandler {
	return &usersHandler{
		config: cfg,
	}
}

func NewUsersRouter(cfg *config.ApiConfig) *chi.Mux {
	usersRouter := chi.NewRouter()
	usersHandler := newUsersHandler(cfg)
	usersRouter.Post("/", usersHandler.createUser)
	usersRouter.Get("/", usersHandler.getUser)

	return usersRouter
}
