package main

import (
	"context"
	"encoding/json"
	"net/http"
	"net/url"
	"time"

	"github.com/Fenroe/shortform/internal/database"
)

/*
Exposed for use in testing
*/
type createURLParams struct {
	ID        *string `json:"id"`
	ExpiredAt *int64  `json:"expired_at"`
	Dest      *string `json:"dest"`
}

/*
Exposed for use in testing
*/
type createURLResponse struct {
	Message string `json:"message"`
	URL     struct {
		ID        string `json:"id"`
		ExpiredAt int64  `json:"expired_at"`
		Dest      string `json:"dest"`
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
			count, err := cfg.DB.CheckForURLWithID(context.Background(), randomID)
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
		count, err := cfg.DB.CheckForURLWithID(context.Background(), *body.ID)
		if err != nil {
			respondWithError(w, http.StatusInternalServerError, "The server encountered an error", err)
			return
		}
		if count > 0 {
			respondWithError(w, http.StatusBadRequest, "This ID is already in use", nil)
			return
		}
	}

	if body.ExpiredAt == nil {
		// Generate an expiry date
		// Set to 30 days for now. TODO: Helper function might improve maintainability
		expiryDate := time.Now().Add(time.Hour * 24 * 30).Unix()
		body.ExpiredAt = &expiryDate
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

	url, err := cfg.DB.CreateURL(
		context.Background(),
		database.CreateURLParams{
			ID:        *body.ID,
			ExpiredAt: time.Unix(*body.ExpiredAt, 0),
			Dest:      *body.Dest,
		},
	)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "The server encountered an error", err)
		return
	}
	var res createURLResponse
	res.Message = "URL created successfully"
	res.URL.Dest = url.Dest
	res.URL.ID = url.ID
	res.URL.ExpiredAt = url.ExpiredAt.Unix()
	respondWithJSON(w, http.StatusCreated, res)
}
