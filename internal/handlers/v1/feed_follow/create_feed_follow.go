package feedfollow

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/Kaivv1/blog-aggregator/internal/database"
	"github.com/Kaivv1/blog-aggregator/pkg/utils"
	"github.com/google/uuid"
)

type FeedFollow struct {
	ID        uuid.UUID `json:"id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	UserID    uuid.UUID `json:"user_id"`
	FeedID    uuid.UUID `json:"feed_id"`
}

func remodelFeedFolow(ff database.FeedsFollow) FeedFollow {
	return FeedFollow{
		ID:        ff.ID,
		CreatedAt: ff.CreatedAt,
		UpdatedAt: ff.UpdatedAt,
		UserID:    ff.UserID,
		FeedID:    ff.FeedID,
	}
}

func (ff *feedFollowHandler) CreateFeedFolow(w http.ResponseWriter, r *http.Request, user database.User) {
	type parameters struct {
		Feed_Id uuid.UUID `json:"feed_id"`
	}
	decoder := json.NewDecoder(r.Body)
	params := parameters{}
	err := decoder.Decode(&params)
	if err != nil {
		utils.RespondWithError(w, http.StatusInternalServerError, "Error decoding paramas in create feed_folow")
		return
	}
	feedFollow, err := ff.config.DB.CreateFeedFollow(r.Context(), database.CreateFeedFollowParams{
		ID:        uuid.New(),
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		FeedID:    params.Feed_Id,
		UserID:    user.ID,
	})
	if err != nil {
		utils.RespondWithError(w, http.StatusInternalServerError, "Error adding feed_follow to db")
		return
	}

	utils.RespondWithJSON(w, http.StatusCreated, remodelFeedFolow(feedFollow))
}
