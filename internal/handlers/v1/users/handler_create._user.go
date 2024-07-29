package users

import (
	"encoding/json"
	"net/http"
	"time"

	db "github.com/Kaivv1/blog-aggregator/internal/database"
	"github.com/Kaivv1/blog-aggregator/pkg/utils"
	"github.com/google/uuid"
)

type User struct {
	ID         string    `json:"id"`
	Created_At time.Time `json:"created_at"`
	Updated_At time.Time `json:"updated_at"`
	Name       string    `json:"name"`
	Api_Key    string    `json:"api_key"`
}

func remodelUser(user db.User) User {
	return User{
		ID:         user.ID.String(),
		Created_At: user.CreatedAt,
		Updated_At: user.UpdatedAt,
		Name:       user.Name,
		Api_Key:    user.ApiKey,
	}
}

func (u *usersHandler) createUser(w http.ResponseWriter, r *http.Request) {
	type parameters struct {
		Name string `json:"name"`
	}

	decoder := json.NewDecoder(r.Body)
	params := parameters{}
	err := decoder.Decode(&params)
	if err != nil {
		utils.RespondWithError(w, http.StatusInternalServerError, "Error while decoding request body at create user")
		return
	}
	user, err := u.config.DB.CreateUser(r.Context(), db.CreateUserParams{
		ID:        uuid.New(),
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		Name:      params.Name,
	})

	if err != nil {
		utils.RespondWithError(w, http.StatusInternalServerError, "Error while creating user in db")
		return
	}

	utils.RespondWithJSON(w, http.StatusCreated, User{
		ID:         user.ID.String(),
		Created_At: user.CreatedAt,
		Updated_At: user.UpdatedAt,
		Name:       user.Name,
		Api_Key:    user.ApiKey,
	})
}
