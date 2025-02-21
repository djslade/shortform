package main

import (
	"net/http"

	"github.com/Fenroe/shortform/internal/database"
)

func handlerGetURLSAsUser(w http.ResponseWriter, r *http.Request) {
	type response struct {
		URLs []database.Url
	}

	wi
}
