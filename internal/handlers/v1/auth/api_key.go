package auth

import (
	"errors"
	"net/http"
	"strings"
)

func GetApiKey(headers http.Header) (string, error) {
	authHeader := headers.Get("Authorization")
	if authHeader == "" {
		return "", errors.New("no auth header provided")
	}
	splittedHeader := strings.Split(authHeader, " ")
	apiKey := splittedHeader[1]
	if apiKey == "" {
		return "", errors.New("no api key")
	}
	return apiKey, nil
}
