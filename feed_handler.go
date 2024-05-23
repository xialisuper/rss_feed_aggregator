package main

import (
	"encoding/json"
	"net/http"

	"github.com/google/uuid"
	"github.com/xialisuper/rss_feed_aggregator/internal/database"
)

func (cfg *apiConfig) handleCreateFeed(w http.ResponseWriter, r *http.Request) {
	body := struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	}{}

	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}

	// get user info from the request context
	user := r.Context().Value(userKey).(database.User)

	// create a new feed in the database
	feed, err := cfg.DB.CreateFeed(r.Context(), database.CreateFeedParams{
		ID:     uuid.New(),
		Name:   body.Name,
		Url:    body.URL,
		UserID: user.ID,
	})

	if err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	// respond with the newly created feed
	respondWithJSON(w, http.StatusCreated, feed)

}
