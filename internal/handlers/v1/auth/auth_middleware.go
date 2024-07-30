package auth

import (
	"net/http"

	"github.com/Kaivv1/blog-aggregator/internal/config"
	"github.com/Kaivv1/blog-aggregator/internal/database"
	"github.com/Kaivv1/blog-aggregator/pkg/utils"
)

type middleware struct {
	config *config.ApiConfig
}

func NewMiddleware(cfg *config.ApiConfig) *middleware {
	return &middleware{
		config: cfg,
	}
}

type nextHandler func(w http.ResponseWriter, r *http.Request, user database.User)

func (m *middleware) AuthMiddleware(handler nextHandler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		apiKey, err := GetApiKey(r.Header)
		if err != nil {
			utils.RespondWithError(w, http.StatusForbidden, err.Error())
			return
		}
		user, err := m.config.DB.GetUserByKey(r.Context(), apiKey)
		if err != nil {
			utils.RespondWithError(w, http.StatusNotFound, "no user matches the key")
			return
		}
		handler(w, r, user)
	}
}
