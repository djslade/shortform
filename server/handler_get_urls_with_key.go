package main

import (
	"net/http"

	"github.com/Fenroe/shortform/internal/database"
)

func handlerGetURLsWithKey(w http.ResponseWriter, r *http.Request) {
	type response struct {
		URLs []database.Url
	}

}
