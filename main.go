package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/joho/godotenv"
)

func main() {
	// Load environment variables from.env file
	godotenv.Load()
	fmt.Println("Starting server...")

	port := os.Getenv("PORT")

	mux := http.NewServeMux()
	server := http.Server{
		Addr:    ":" + port,
		Handler: mux,
	}

	// GET /v1/readiness
	mux.HandleFunc("GET /v1/readiness", handleReadinessProbe)
	mux.HandleFunc("GET /v1/err", errorHandler)

	fmt.Println("Server running on port ", port)

	err := server.ListenAndServe()

	if err != nil {
		panic(err)
	}

}


