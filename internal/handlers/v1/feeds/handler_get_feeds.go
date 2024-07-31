package feeds

import (
	"net/http"

	"github.com/Kaivv1/blog-aggregator/pkg/utils"
)

func (f *feedsHandler) GetFeeds(w http.ResponseWriter, r *http.Request) {
	feeds, err := f.config.DB.GetFeeds(r.Context())
	if err != nil {
		utils.RespondWithJSON(w, http.StatusInternalServerError, "Error getting feeds from db")
		return
	}
	remodeledFeeds := make([]Feed, len(feeds))

	for i, feed := range feeds {
		remodeledFeeds[i] = Feed{
			ID:        feed.ID,
			CreatedAt: feed.CreatedAt,
			UpdatedAt: feed.UpdatedAt,
			Name:      feed.Name,
			Url:       feed.Url,
			UserID:    feed.UserID,
		}
	}

	utils.RespondWithJSON(w, http.StatusOK, remodeledFeeds)
}
