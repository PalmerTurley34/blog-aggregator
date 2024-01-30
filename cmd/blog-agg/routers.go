package main

import (
	"github.com/go-chi/chi/v5"
)

func newV1Router(cfg *apiConfig) *chi.Mux {
	router := chi.NewRouter()
	router.Get("/readiness", getReadiness)
	router.Get("/err", getErr)
	router.Post("/users", cfg.createUser)
	router.Get("/users", cfg.middlewareAuth(cfg.getUserByApiKey))

	router.Post("/feeds", cfg.middlewareAuth(cfg.createFeed))
	router.Get("/feeds", cfg.getAllFeeds)

	return router
}
