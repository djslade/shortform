package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
)

func main() {
	// Retrieve env variables
	godotenv.Load(".env")

	port := os.Getenv("PORT")
	if port == "" {
		log.Fatal("PORT env variables is not set")
	}

	// Handler
	mux := http.NewServeMux()

	mux.HandleFunc("GET /url", handlerURL)
	// Server
	server := http.Server{
		Addr:    ":" + port,
		Handler: mux,
	}

	// Start server
	fmt.Printf("Starting server on http://localhost:%s\n", port)
	log.Fatal(server.ListenAndServe())
}
