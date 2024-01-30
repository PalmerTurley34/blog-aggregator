package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/PalmerTurley34/blog-aggregator/internal/database"
	"github.com/PalmerTurley34/blog-aggregator/internal/models"
	"github.com/google/uuid"
)

func (cfg *apiConfig) createFeedFollow(w http.ResponseWriter, r *http.Request, user database.User) {
	type parameters struct {
		FeedID uuid.UUID `json:"feed_id"`
	}
	params := parameters{}
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&params)
	if err != nil {
		respondWithError(w, http.StatusBadRequest, fmt.Sprintf("couldn't parse request body: %v", err))
		return
	}

	newFeedFollow, err := cfg.DB.CreateFeedFollow(r.Context(), database.CreateFeedFollowParams{
		ID:        uuid.New(),
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		UserID:    user.ID,
		FeedID:    params.FeedID,
	})

	if err != nil {
		respondWithError(w, http.StatusInternalServerError, fmt.Sprintf("Couldn't create feed_follow: %v", err))
		return
	}
	respondWithJSON(w, http.StatusOK, models.DBFeedFollowConvert(newFeedFollow))
}

func (cfg *apiConfig) getUsersFeedFollows(w http.ResponseWriter, r *http.Request, user database.User) {
	feedFollows, err := cfg.DB.GetFeedFollows(r.Context(), user.ID)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, fmt.Sprintf("couldn't get feed follows: %v", err))
		return
	}
	respondWithJSON(w, http.StatusOK, models.DBFeedFollowsConvert(feedFollows))
}
