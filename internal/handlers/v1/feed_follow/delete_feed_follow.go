package feedfollow

import (
	"net/http"

	"github.com/Kaivv1/blog-aggregator/internal/database"
	"github.com/Kaivv1/blog-aggregator/pkg/utils"
	"github.com/go-chi/chi"
	"github.com/google/uuid"
)

func (ff *feedFollowHandler) DeleteFeedFollow(w http.ResponseWriter, r *http.Request, user database.User) {
	val := chi.URLParam(r, "ff_id")
	if val == "" {
		utils.RespondWithError(w, http.StatusBadRequest, "Cant find feed_follow id")
		return
	}
	feedFollowId, err := uuid.Parse(val)
	if err != nil {
		utils.RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	err = ff.config.DB.DeleteFeedFollow(r.Context(), feedFollowId)
	if err != nil {
		utils.RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
}
