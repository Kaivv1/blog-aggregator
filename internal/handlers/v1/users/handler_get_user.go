package users

import (
	"net/http"

	"github.com/Kaivv1/blog-aggregator/internal/database"
	"github.com/Kaivv1/blog-aggregator/pkg/utils"
)

func (u *usersHandler) GetUser(w http.ResponseWriter, r *http.Request, user database.User) {
	utils.RespondWithJSON(w, http.StatusOK, RemodelUser(user))
}
