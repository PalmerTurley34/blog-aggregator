package routes

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/PalmerTurley34/blog-aggregator/internal/database"
	"github.com/PalmerTurley34/blog-aggregator/internal/responses"
	"github.com/go-chi/chi/v5"
	"github.com/google/uuid"
)

type ApiConfig struct {
	DB *database.Queries
}

func NewV1Router(cfg *ApiConfig) *chi.Mux {
	router := chi.NewRouter()
	router.Get("/readiness", getReadiness)
	router.Get("/err", getErr)
	router.Post("/users", cfg.createUser)

	return router
}

func getReadiness(w http.ResponseWriter, r *http.Request) {
	type returnBody struct {
		Status string `json:"status"`
	}
	responses.WithJSON(w, http.StatusOK, returnBody{Status: "ok"})
}

func getErr(w http.ResponseWriter, r *http.Request) {
	responses.WithError(w, http.StatusInternalServerError, "Internal Server Error")
}

func (cfg *ApiConfig) createUser(w http.ResponseWriter, r *http.Request) {
	type parameters struct {
		Name string `json:"name"`
	}
	params := parameters{}
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&params)
	if err != nil {
		responses.WithError(w, http.StatusBadRequest, fmt.Sprintf("Cannot decode JSON body: %v", err.Error()))
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
			fmt.Sprintf("Error creating new user: %v", err.Error()))
		return
	}
	responses.WithJSON(w, http.StatusCreated, newUser)
}
