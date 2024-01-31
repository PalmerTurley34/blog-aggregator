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

func (cfg *apiConfig) createUser(w http.ResponseWriter, r *http.Request) {
	type parameters struct {
		Name string `json:"name"`
	}
	params := parameters{}
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&params)
	if err != nil {
		respondWithError(w, http.StatusBadRequest, fmt.Sprintf("Cannot decode JSON body: %v", err))
		return
	}
	newUser, err := cfg.DB.CreateUser(r.Context(), database.CreateUserParams{
		ID:        uuid.New(),
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		Name:      params.Name,
	})
	if err != nil {
		respondWithError(
			w,
			http.StatusInternalServerError,
			fmt.Sprintf("Error creating new user: %v", err))
		return
	}
	respondWithJSON(w, http.StatusCreated, models.DBUserConvert(newUser))
}

func (cfg *apiConfig) getUserByApiKey(w http.ResponseWriter, r *http.Request, user database.User) {
	respondWithJSON(w, http.StatusOK, models.DBUserConvert(user))
}

func (cfg *apiConfig) getPostsForUser(w http.ResponseWriter, r *http.Request, user database.User) {
	posts, err := cfg.DB.GetPostsByUser(r.Context(), database.GetPostsByUserParams{
		UserID: user.ID,
		Limit:  5,
	})
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, fmt.Sprintf("Couldn't get posts: %v", err))
		return
	}

	respondWithJSON(w, http.StatusOK, models.DBPostsConvert(posts))
}
