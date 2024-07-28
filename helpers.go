package main

import (
	"encoding/json"
	"log"
	"net/http"
)

func respondWithJSON(w http.ResponseWriter, code int, payload interface{}) {
	w.Header().Add("Content-Type:", "application/json")
	data, err := json.Marshal(payload)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Printf("Cannot marshal data: %s\n", err)
		return
	}
	w.WriteHeader(code)
	w.Write(data)
}

func respondWithError(w http.ResponseWriter, code int, msg string) {
	type respondWithError struct {
		Error string `json:"error"`
	}
	if code > 499 {
		log.Printf("Failed with status code 5XX: %s\n", msg)
	}
	respondWithJSON(w, code, respondWithError{
		Error: msg,
	})
}
