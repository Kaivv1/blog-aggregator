package users

import (
	"fmt"
	"net/http"

	"github.com/Kaivv1/blog-aggregator/internal/handlers/v1/auth"
	"github.com/Kaivv1/blog-aggregator/pkg/utils"
)

func (u *usersHandler) getUser(w http.ResponseWriter, r *http.Request) {
	apiKey, err := auth.GetApiKey(r.Header)
	if err != nil {
		utils.RespondWithError(w, http.StatusForbidden, fmt.Sprintf("Auth error: %v", err))
		return
	}
	user, err := u.config.DB.GetUserByKey(r.Context(), apiKey)
	if err != nil {
		utils.RespondWithError(w, http.StatusNotFound, "no user matches the key")
		return
	}

	utils.RespondWithJSON(w, http.StatusOK, remodelUser(user))
}
