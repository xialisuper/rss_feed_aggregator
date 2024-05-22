package main

import (
	"encoding/json"
	"net/http"

	"github.com/google/uuid"
	"github.com/xialisuper/rss_feed_aggregator/internal/database"
)

func (cgf *apiConfig) handleCreateUser(w http.ResponseWriter, r *http.Request) {
	// get user data from request body
	u := struct {
		Name string `json:"name"`
	}{}
	err := json.NewDecoder(r.Body).Decode(&u)
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid request payload")
	}

	// create user in database
	ctx := r.Context()
	newUser, err := cgf.DB.CreateUser(ctx, database.CreateUserParams{ID: uuid.New(), Name: u.Name})
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Failed to create user")
		return
	}

	// respond with success message
	respondWithJSON(w, http.StatusCreated, newUser)

}
