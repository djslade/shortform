package main

import (
	"context"
	"net/http"
	"time"
)

type redirectResponse struct {
	Message  string `json:"message"`
	Redirect string `json:"redirect"`
}

func (cfg *apiConfig) handlerRedirect(w http.ResponseWriter, r *http.Request) {
	urlID := r.URL.Query().Get("url")
	if urlID == "" {
		// TODO: Bad request
		respondWithError(w, http.StatusBadRequest, "URL query parameter is missing", nil)
		return
	}
	url, err := cfg.db.GetURLByID(context.Background(), urlID)
	if err != nil {
		// TODO: Not found
		respondWithError(w, http.StatusNotFound, "URL not found", nil)
		return
	}
	// Check if url is disabled
	if url.ExpiredAt.Time.Unix() <= time.Now().Unix() {
		// TODO: Bad request
		respondWithError(w, http.StatusBadRequest, "URL has expired", nil)
		return
	}
	var res redirectResponse
	res.Message = "Redirect to supplied URL"
	res.Redirect = url.Destination
	// Redirect
	respondWithJSON(w, http.StatusFound, res)
}
