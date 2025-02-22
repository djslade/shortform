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
	db             database.Queries
	jwtSecret      string
	localIpAddress string
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

	localIpAddress := os.Getenv("LOCAL_IP_ADDRESS")
	if localIpAddress == "" {
		log.Fatal("LOCAL_IP_ADDRESS env variable is not set")
	}

	// DB setup
	db, err := sql.Open("postgres", dbConnectionString)
	if err != nil {
		log.Fatal("Unable to start database")
	}

	// Config
	config := apiConfig{
		db:             *database.New(db),
		jwtSecret:      jwtSecret,
		localIpAddress: localIpAddress,
	}

	// Handler
	mux := http.NewServeMux()

	mux.HandleFunc("POST /api/key", config.handlerCreateKey)

	mux.HandleFunc("GET /api/url-key", config.handlerGetURLsWithKey)
	mux.HandleFunc("POST /api/url-key", config.handlerCreateURLWithKey)
	mux.HandleFunc("POST /api/url-auth", config.handlerCreateURLAsUser)

	mux.HandleFunc("POST /api/users", config.handlerCreateUser)

	mux.HandleFunc("POST /api/login", config.handlerLogin)
	mux.HandleFunc("POST /api/refresh", config.handlerRefresh)
	mux.HandleFunc("DELETE /api/revoke", config.handlerRevoke)

	mux.HandleFunc("POST /api/redirect/{urlID}", config.handlerRedirect)

	mux.HandleFunc("GET /{urlID}", config.handlerRedirect)

	// Server
	server := http.Server{
		Addr:    ":" + port,
		Handler: mux,
	}

	// Start server
	fmt.Printf("Starting server on http://localhost:%s\n", port)
	log.Fatal(server.ListenAndServe())
}
