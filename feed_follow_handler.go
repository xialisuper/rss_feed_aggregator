package main

import (
	"encoding/json"
	"net/http"

	"github.com/google/uuid"
	"github.com/xialisuper/rss_feed_aggregator/internal/database"
)

func (cfg *apiConfig) handleCreateFeedFollow(w http.ResponseWriter, r *http.Request) {
	body := struct {
		FeedID string `json:"feed_id"`
	}{}

	// parse request body
	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Convert string to UUID
	feedIDUUID, err := uuid.Parse(body.FeedID)
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid feed ID")
		return
	}
	// check if feed exists
	_, err = cfg.DB.GetFeedById(r.Context(), feedIDUUID)
	if err != nil {
		respondWithError(w, http.StatusNotFound, "Feed not found")
		return
	}

	user := r.Context().Value(userKey).(database.User)

	// create follow
	newFollow, err := cfg.DB.CreateFeedFollow(r.Context(), database.CreateFeedFollowParams{
		ID:     uuid.New(),
		FeedID: feedIDUUID,
		UserID: user.ID,
	})

	if err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	// return response
	respondWithJSON(w, http.StatusCreated, newFollow)
}

func (cfg *apiConfig) handleDeleteFeedFollow(w http.ResponseWriter, r *http.Request) {

	feedFollowID := r.PathValue("feedFollowID")
	// Convert string to UUID
	feedFollowIDUUID, err := uuid.Parse(feedFollowID)
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid feed follow ID")
		return
	}

	// delete follow by user_id and feed_id. they should match.
	err = cfg.DB.DeleteFeedFollowByID(r.Context(), database.DeleteFeedFollowByIDParams{
		UserID: r.Context().Value(userKey).(database.User).ID,
		FeedID: feedFollowIDUUID,
	})
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	respondWithJSON(w, http.StatusNoContent, nil)

}
