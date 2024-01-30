package main

import (
	"fmt"
	"net/http"

	"github.com/PalmerTurley34/blog-aggregator/internal/auth"
	"github.com/PalmerTurley34/blog-aggregator/internal/database"
	"github.com/PalmerTurley34/blog-aggregator/internal/responses"
)

type authedHandler func(http.ResponseWriter, *http.Request, database.User)

func (cfg *apiConfig) middlewareAuth(handler authedHandler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		apiKey, err := auth.GetApiKey(r.Header)
		if err != nil {
			responses.WithError(w, http.StatusForbidden, fmt.Sprintf("Error with ApiKey: %v", err))
			return
		}
		user, err := cfg.DB.GetUserByApiKey(r.Context(), apiKey)
		if err != nil {
			responses.WithError(w, http.StatusBadRequest, fmt.Sprintf("Error with DB: %v", err))
			return
		}
		handler(w, r, user)
	}
}
