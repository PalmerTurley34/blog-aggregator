package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/PalmerTurley34/blog-aggregator/internal/database"
	"github.com/PalmerTurley34/blog-aggregator/internal/models"
	"github.com/PalmerTurley34/blog-aggregator/internal/responses"
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
		responses.WithError(w, http.StatusBadRequest, fmt.Sprintf("Cannot decode JSON body: %v", err))
	}
	newUser, err := cfg.DB.CreateUser(r.Context(), database.CreateUserParams{
		ID:        uuid.New(),
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		Name:      params.Name,
	})
	if err != nil {
		responses.WithError(
			w,
			http.StatusInternalServerError,
			fmt.Sprintf("Error creating new user: %v", err))
		return
	}
	responses.WithJSON(w, http.StatusCreated, models.DBUserToUser(newUser))
}

func getUserByApiKey(w http.ResponseWriter, r *http.Request, user database.User) {
	responses.WithJSON(w, http.StatusOK, models.DBUserToUser(user))
}
