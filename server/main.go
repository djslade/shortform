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
	db        database.Queries
	jwtSecret string
}

func main() {
	// Retrieve env variables
	godotenv.Load(".env")

	port := os.Getenv("PORT")
	if port == "" {
		log.Fatal("PORT env variable is not set")
	}

	dbConnectionString := os.Getenv("PG_CONNECTION_STRING")
	if dbConnectionString == "" {
		log.Fatal("PG_CONNECTION_STRING env variable is not set")
	}

	jwtSecret := os.Getenv("JWT_SECRET")
	if jwtSecret == "" {
		log.Fatal("JWT_SECRET env variable is not set")
	}

	// DB setup
	db, err := sql.Open("postgres", dbConnectionString)
	if err != nil {
		log.Fatal("Unable to start database")
	}

	// Config
	config := apiConfig{
		db:        *database.New(db),
		jwtSecret: jwtSecret,
	}

	// Handler
	mux := http.NewServeMux()

	mux.HandleFunc("GET /url", config.handlerURL)
	mux.HandleFunc("GET /redirect", config.handlerRedirect)
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
