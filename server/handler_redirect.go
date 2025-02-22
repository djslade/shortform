package main

import (
	"context"
	"database/sql"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"

	"github.com/Fenroe/shortform/internal/database"
)

func (cfg *apiConfig) handlerRedirect(w http.ResponseWriter, r *http.Request) {
	urlID := r.PathValue("urlID")
	if urlID == "" {
		respondWithError(w, http.StatusBadRequest, "No URL was specified", errors.New("no URL ID"))
		return
	}

	url, err := cfg.db.GetURLByID(context.Background(), urlID)
	if err != nil {
		respondWithError(w, http.StatusNotFound, "URL was not found", err)
		return
	}
	http.Redirect(w, r, url.Destination, http.StatusFound)

	go func() {
		type ipapiResponse struct {
			Status     string  `json:"status"`
			Continent  string  `json:"continent"`
			Country    string  `json:"country"`
			RegionName string  `json:"regionName"`
			City       string  `json:"city"`
			Lat        float64 `json:"lat"`
			Lon        float64 `json:"lon"`
			Timezone   string  `json:"timezone"`
			Currency   string  `json:"currency"`
			Isp        string  `json:"isp"`
			Mobile     bool    `json:"mobile"`
			Proxy      bool    `json:"proxy"`
		}

		ipAddress := getClientIP(r)
		if ipAddress == "::1" {
			// The request came from the same machine as the server is running on
			// This only happens in testing
			ipAddress = cfg.localIpAddress
		}
		res, err := http.Get(fmt.Sprintf("http://ip-api.com/json/%s?fields=9683929", ipAddress))
		if err != nil {
			log.Println(err)
			return
		}
		defer res.Body.Close()
		body, err := io.ReadAll(res.Body)
		if err != nil {
			log.Println(err)
			return
		}
		var apiRes ipapiResponse
		if err = json.Unmarshal(body, &apiRes); err != nil {
			log.Println(err)
			return
		}
		if apiRes.Status == "fail" {
			log.Println(apiRes.Status)
			return
		}
		referralURL := r.Header.Get("referer")
		log.Println(referralURL)
		click, err := cfg.db.CreateClick(context.Background(), database.CreateClickParams{
			UrlID:     getNullString(url.ID), // TODO: urlID should never be null
			Continent: getNullString(apiRes.Continent),
			Country:   getNullString(apiRes.Country),
			Region:    getNullString(apiRes.RegionName),
			City:      getNullString(apiRes.City),
			Lat: sql.NullFloat64{
				Valid:   true,
				Float64: apiRes.Lat,
			},
			Lon: sql.NullFloat64{
				Valid:   true,
				Float64: apiRes.Lon,
			},
			Timezone:    getNullString(apiRes.Timezone),
			Currency:    getNullString(apiRes.Currency),
			ReferralUrl: getNullString(referralURL),
			IsMobile: sql.NullBool{
				Valid: true,
				Bool:  apiRes.Mobile,
			},
			IsProxy: sql.NullBool{
				Valid: true,
				Bool:  apiRes.Proxy,
			},
			Isp: getNullString(apiRes.Isp),
		})
		if err != nil {
			log.Println(err)
			return
		}
		log.Println(click.ID)
	}()
}
