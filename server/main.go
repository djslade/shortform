package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/Fenroe/shortform/internal/database"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

type apiConfig struct {
	DB *database.Queries
}

func main() {
	// Retrieve env variables
	godotenv.Load(".env")

	port := os.Getenv("PORT")
	if port == "" {
		log.Fatal("PORT env variable is not set")
	}

	dbConnectionString := os.Getenv("DB_CONNECTION_STRING")
	if dbConnectionString == "" {
		log.Fatal("DB_CONNECTION_STRING env variable is not set")
	}

	// DB setup
	db, err := sql.Open("postgres", dbConnectionString)
	if err != nil {
		log.Fatal("Unable to start database")
	}

	// Config
	config := apiConfig{
		DB: database.New(db),
	}

	// Handler
	mux := http.NewServeMux()

	mux.HandleFunc("GET /url", config.handlerURL)
	mux.HandleFunc("POST /url", config.handlerCreateURL)

	// Server
	server := http.Server{
		Addr:    ":" + port,
		Handler: mux,
	}

	// Start server
	fmt.Printf("Starting server on http://localhost:%s\n", port)
	log.Fatal(server.ListenAndServe())
}
