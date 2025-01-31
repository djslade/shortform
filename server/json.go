package main

import (
	"encoding/json"
	"log"
	"net/http"
)

/*
Middleware for responding with errors in API endpoints.

Logs the error itself to the console and forwards a client-friendly response struct
to the respondWithJson method. If the code indicates an Internal Server Eror, that is
emphasized in the console.
*/
func respondWithError(w http.ResponseWriter, code int, msg string, err error) {
	if err != nil {
		log.Println(err)
	}
	if code > 499 {
		log.Printf("Responding with 5XX error: %s", msg)
	}
	type errorResponse struct {
		Message string `json:"message"`
	}
	respondWithJSON(w, code, errorResponse{
		Message: msg,
	})
}

/*
General-purpose middleware for handling responses in API endpoints.

Sets JSON header, statusCode and writes data to response body in JSON format.
*/
func respondWithJSON(w http.ResponseWriter, code int, payload interface{}) {
	w.Header().Set("Content-Type", "application/json")
	data, err := json.Marshal(payload)
	if err != nil {
		log.Printf("Error marshalling JSON: %s", err)
		w.WriteHeader(500)
		return
	}
	w.WriteHeader(code)
	w.Write(data)
}
