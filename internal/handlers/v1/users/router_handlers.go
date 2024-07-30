package users

import (
	"github.com/Kaivv1/blog-aggregator/internal/config"
	"github.com/Kaivv1/blog-aggregator/internal/handlers/v1/auth"
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
	m := auth.NewMiddleware(cfg)
	usersRouter := chi.NewRouter()
	usersHandler := newUsersHandler(cfg)
	usersRouter.Post("/", usersHandler.CreateUser)
	usersRouter.Get("/", m.AuthMiddleware(usersHandler.GetUser))

	return usersRouter
}
