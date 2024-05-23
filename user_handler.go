package main

import (
	"encoding/json"
	"errors"
	"net/http"
	"strings"

	"github.com/google/uuid"
	"github.com/xialisuper/rss_feed_aggregator/internal/database"
)


// handleUpdateUser updates a user in the database
func (cgf *apiConfig) handleUpdateUserByApiKey(w http.ResponseWriter, r *http.Request) {
	// get user data from request body
	u := struct {
		Name string `json:"name"`
	}{}
	err := json.NewDecoder(r.Body).Decode(&u)
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}

	// get user from database
	ctx := r.Context()
	apiKey, err := getApiKeyFromHeader(r)
	if err != nil {
		respondWithError(w, http.StatusUnauthorized, "Invalid api key")
		return
	}

	// update user in database
	updatedUser, err := cgf.DB.UpdateUser(ctx, database.UpdateUserParams{ApiKey: apiKey, Name: u.Name})
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Failed to update user")
		return
	}

	// respond with success message
	respondWithJSON(w, http.StatusOK, updatedUser)

}

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

// handleGetUsersByApiKey returns a user by api key
func (cgf *apiConfig) handleGetUsersByApiKey(w http.ResponseWriter, r *http.Request) {

	apiKey, err := getApiKeyFromHeader(r)
	if err != nil {
		respondWithError(w, http.StatusUnauthorized, "Invalid api key")
		return
	}

	user, err := cgf.DB.GetUserByApiKey(r.Context(), apiKey)
	if err != nil {
		respondWithError(w, http.StatusNotFound, "User not found")
		return
	}

	respondWithJSON(w, http.StatusOK, user)

}

func getApiKeyFromHeader(r *http.Request) (string, error) {
	authHeader := r.Header.Get("Authorization")
	keyParts := strings.Split(authHeader, " ")
	var apiKey string

	if len(keyParts) == 2 && keyParts[0] == "ApiKey" {
		apiKey = keyParts[1]
		return apiKey, nil

	}

	return "", errors.New("invalid api key")
}
