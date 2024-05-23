package main

import (
	"database/sql"
	"fmt"
	"net/http"
	"os"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"github.com/xialisuper/rss_feed_aggregator/internal/database"
)

func main() {
	// Load environment variables from.env file
	godotenv.Load()
	fmt.Println("Starting server...")

	// start db connection here
	path := os.Getenv("DB_CONN_STR")
	db, err := sql.Open("postgres", path)

	if err != nil {
		panic(err)
	}

	fmt.Println("Connected to database")
	defer db.Close()

	dbQueries := database.New(db)

	apiConfig := apiConfig{
		DB: dbQueries,
	}

	port := os.Getenv("PORT")

	mux := http.NewServeMux()
	server := http.Server{
		Addr:    ":" + port,
		Handler: mux,
	}

	// GET /v1/readiness
	mux.HandleFunc("GET /v1/readiness", handleReadinessProbe)
	mux.HandleFunc("GET /v1/err", errorHandler)
	// POST /v1/users
	mux.HandleFunc("POST /v1/users", apiConfig.handleCreateUser)
	// GET /v1/users
	mux.HandleFunc("GET /v1/users", apiConfig.handleGetUsersByApiKey)
	// GET /v1/users/{id}
	// mux.HandleFunc("GET /v1/users/{id}", apiConfig.handleGetUser)
	// PUT /v1/users
	mux.HandleFunc("PUT /v1/users", apiConfig.handleUpdateUserByApiKey)

	fmt.Println("Server running on port ", port)

	err = server.ListenAndServe()

	if err != nil {
		panic(err)
	}

}
