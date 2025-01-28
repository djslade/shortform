package main

import (
	"context"
	"encoding/json"
	"net/http"
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
		// TODO: Internal Server Error
		return
	}

	if body.ID == nil {
		// TODO: Bad request
		return
	}

	if body.ExpiredAt == nil {
		// TODO: Bad request
		return
	}

	if body.Dest == nil {
		// TODO: Bad request
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
		// TODO: Handle database error
		return
	}
	var res createURLResponse
	res.Message = "URL created successfully"
	res.URL.Dest = url.Dest
	res.URL.ID = url.ID
	res.URL.ExpiredAt = url.ExpiredAt.Unix()
	data, err := json.Marshal(res)
	if err != nil {
		// TODO: Internal Server error
		return
	}
	w.WriteHeader(http.StatusCreated)
	w.Write(data)
}
