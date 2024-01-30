package main

import (
	"github.com/go-chi/chi/v5"
)

func newV1Router(cfg *apiConfig) *chi.Mux {
	router := chi.NewRouter()
	router.Get("/readiness", getReadiness)
	router.Get("/err", getErr)
	router.Post("/users", cfg.createUser)
	router.Get("/users", cfg.middlewareAuth(getUserByApiKey))

	return router
}
