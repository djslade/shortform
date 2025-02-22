package main

import (
	"context"
	"net/http"

	"github.com/Fenroe/shortform/internal/auth"
)

func (cfg *apiConfig) handlerCreateKey(w http.ResponseWriter, r *http.Request) {
	type response struct {
		Key string `json:"key"`
	}

	newKey, err := auth.MakeAPIKey()
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Could not generate new key", err)
		return
	}
	apiKey, err := cfg.db.CreateAPIKey(context.Background(), newKey)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Could not create key", err)
		return
	}
	respondWithJSON(w, http.StatusCreated, response{Key: apiKey.Key})
}
