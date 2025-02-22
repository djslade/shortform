package main

import (
	"context"
	"database/sql"
	"net/http"

	"github.com/Fenroe/shortform/internal/auth"
	"github.com/Fenroe/shortform/internal/database"
)

func (cfg *apiConfig) handlerGetURLsWithKey(w http.ResponseWriter, r *http.Request) {
	type response struct {
		URLs []database.Url
	}

	key, err := auth.GetAPIKey(r.Header)
	if err != nil {
		respondWithError(w, http.StatusUnauthorized, "API key not found in header", err)
		return
	}
	if _, err = cfg.db.GetAPIKey(context.Background(), key); err != nil {
		respondWithError(w, http.StatusNotFound, "Could not find API key", err)
	}
	urls, err := cfg.db.GetURLsByAPIKey(context.Background(), sql.NullString{
		Valid:  true,
		String: key,
	})
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Could not execute query", err)
	}
	respondWithJSON(w, http.StatusOK, response{URLs: urls})
}
