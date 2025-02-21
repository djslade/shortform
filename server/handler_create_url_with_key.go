package main

import (
	"context"
	"database/sql"
	"encoding/json"
	"net/http"
	"net/url"

	"github.com/Fenroe/shortform/internal/auth"
	"github.com/Fenroe/shortform/internal/database"
)

func (cfg *apiConfig) handlerCreateURLWithKey(w http.ResponseWriter, r *http.Request) {
	type parameters struct {
		Dest *string `json:"dest"`
	}

	type response struct {
		URL database.Url
	}

	key, err := auth.GetAPIKey(r.Header)
	if err != nil {
		respondWithError(w, http.StatusUnauthorized, "API key not found in header", err)
		return
	}
	if _, err = cfg.db.GetAPIKey(context.Background(), key); err != nil {
		respondWithError(w, http.StatusNotFound, "Could not find API key", err)
		return
	}

	urlCount, err := cfg.db.GetURLsByAPIKeyCount(context.Background(), sql.NullString{
		Valid:  true,
		String: key,
	})
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Could not execute query", err)
		return
	}
	if urlCount >= 3 {
		respondWithError(w, http.StatusUnauthorized, "No more URLs can be created with this key", err)
		return
	}

	var body parameters
	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		respondWithError(w, http.StatusInternalServerError, "The server encountered an error", err)
		return
	}

	var id string
	for {
		randomID, err := generateURLID(5)
		if err != nil {
			respondWithError(w, http.StatusInternalServerError, "The server encountered an error", err)
			return
		}
		count, err := cfg.db.CheckForURLWithID(context.Background(), randomID)
		if err != nil {
			respondWithError(w, http.StatusInternalServerError, "The server encountered an error", err)
			return
		}
		if count == 0 {
			id = randomID
			break
		}
	}

	// Dest must be an absolute URL.
	if _, err := url.ParseRequestURI(*body.Dest); err != nil {
		respondWithError(w, http.StatusBadRequest, "Dest must be a valid, absolute URL", nil)
		return
	}

	url, err := cfg.db.CreateURL(
		context.Background(),
		database.CreateURLParams{
			ID:          id,
			Destination: *body.Dest,
		},
	)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Could not create URL", err)
		return
	}
	respondWithJSON(w, http.StatusCreated, response{URL: url})
}
