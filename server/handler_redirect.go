package main

import (
	"context"
	"encoding/json"
	"net/http"
)

func (cfg *apiConfig) handlerRedirect(w http.ResponseWriter, r *http.Request) {
	type parameters struct {
		URLID *string `json:"url_id"`
	}

	var body parameters
	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		respondWithError(w, http.StatusInternalServerError, "The server encountered an error", err)
		return
	}

	url, err := cfg.db.GetURLByID(context.Background(), *body.URLID)
	if err != nil {
		respondWithError(w, http.StatusNotFound, "URL was not found", err)
	}
	http.Redirect(w, r, url.Destination, http.StatusFound)
	go func() {}()
}
