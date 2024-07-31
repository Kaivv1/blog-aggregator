package feedfollow

import (
	"net/http"

	"github.com/Kaivv1/blog-aggregator/internal/database"
	"github.com/Kaivv1/blog-aggregator/pkg/utils"
)

func (ff *feedFollowHandler) GetFollowedFeedsUser(w http.ResponseWriter, r *http.Request, user database.User) {
	followedFeeds, err := ff.config.DB.GetFollowFeedsUser(r.Context(), user.ID)
	if err != nil {
		utils.RespondWithError(w, http.StatusNotFound, err.Error())
		return
	}

	remodel := make([]FeedFollow, len(followedFeeds))

	for i, followedFeed := range followedFeeds {
		remodel[i] = FeedFollow{
			ID:        followedFeed.ID,
			CreatedAt: followedFeed.CreatedAt,
			UpdatedAt: followedFeed.UpdatedAt,
			UserID:    followedFeed.UserID,
			FeedID:    followedFeed.FeedID,
		}
	}

	utils.RespondWithJSON(w, http.StatusOK, remodel)
}
