package main

import (
	"context"
	"fmt"
	"net/http"
)

func (cfg *apiConfig) handlerURL(w http.ResponseWriter, r *http.Request) {
	queires := r.URL.Query()
	queryURL := queires.Get("url")
	if queryURL == "" {
		// TODO: Handle invalid request
		return
	}
	url, err := cfg.db.GetURLByID(context.Background(), queryURL)
	if err != nil {
		// TODO: Handle server error
		return
	}
	fmt.Printf("%v\n", url.ID)
}
