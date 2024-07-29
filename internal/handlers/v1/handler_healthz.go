package v1

import (
	"net/http"

	"github.com/Kaivv1/blog-aggregator/pkg/utils"
)

func HandleHealthz(w http.ResponseWriter, r *http.Request) {
	utils.RespondWithJSON(w, http.StatusOK, struct {
		Status string `json:"status"`
	}{
		Status: "ok",
	})
}
