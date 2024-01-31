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
	router.Get("/users/posts", cfg.middlewareAuth(cfg.getPostsForUser))

	router.Post("/feeds", cfg.middlewareAuth(cfg.createFeed))
	router.Get("/feeds", cfg.getAllFeeds)

	router.Post("/feed_follows", cfg.middlewareAuth(cfg.createFeedFollow))
	router.Get("/feed_follows", cfg.middlewareAuth(cfg.getUsersFeedFollows))
	router.Delete("/feed_follows/{feedFollowID}", cfg.middlewareAuth(cfg.deleteFeedFollow))

	return router
}
