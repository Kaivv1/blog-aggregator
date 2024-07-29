package v1

import (
	"net/http"

	"github.com/Kaivv1/blog-aggregator/pkg/utils"
)

func HandleErr(w http.ResponseWriter, r *http.Request) {
	utils.RespondWithError(w, http.StatusInternalServerError, "Internal server error")
}
