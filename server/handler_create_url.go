package main

import (
	"context"
	"encoding/json"
	"net/http"
	"net/url"

	"github.com/Fenroe/shortform/internal/database"
)

// Exposed for use in testing
type createURLParams struct {
	ID   *string `json:"id"`
	Dest *string `json:"dest"`
}

// Exposed for use in testing
type createURLResponse struct {
	URL struct {
		ID   string `json:"id"`
		Dest string `json:"dest"`
	} `json:"url"`
}

func (cfg *apiConfig) handlerCreateURL(w http.ResponseWriter, r *http.Request) {

	var body createURLParams

	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		respondWithError(w, http.StatusInternalServerError, "The server encountered an error", err)
		return
	}

	if body.ID == nil {
		// TODO: Random ID
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
				body.ID = &randomID
				break
			}
		}

	} else {
		count, err := cfg.db.CheckForURLWithID(context.Background(), *body.ID)
		if err != nil {
			respondWithError(w, http.StatusInternalServerError, "The server encountered an error", err)
			return
		}
		if count > 0 {
			respondWithError(w, http.StatusBadRequest, "This ID is already in use", nil)
			return
		}
	}

	if body.Dest == nil {
		respondWithError(w, http.StatusBadRequest, "Dest field missing from request", nil)
		return
	}

	// Dest must be an absolute URL.
	if _, err := url.ParseRequestURI(*body.Dest); err != nil {
		respondWithError(w, http.StatusBadRequest, "Dest must be a valid, absolute URL", nil)
		return
	}

	url, err := cfg.db.CreateURL(
		context.Background(),
		database.CreateURLParams{
			ID:          *body.ID,
			Destination: *body.Dest,
		},
	)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "The server encountered an error", err)
		return
	}
	var res createURLResponse
	res.URL.Dest = url.Destination
	res.URL.ID = url.ID
	respondWithJSON(w, http.StatusCreated, res)
}
