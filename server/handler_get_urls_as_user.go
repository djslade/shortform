package main

import (
	"net/http"

	"github.com/Fenroe/shortform/internal/database"
)

func (cfg *apiConfig) handlerGetURLsAsUser(w http.ResponseWriter, r *http.Request) {
	type response struct {
		URLs []database.Url
	}

}
