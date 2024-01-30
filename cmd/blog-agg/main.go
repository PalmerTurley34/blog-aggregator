package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/PalmerTurley34/blog-aggregator/internal/database"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/cors"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

type apiConfig struct {
	DB *database.Queries
}

func main() {
	godotenv.Load()

	port := os.Getenv("PORT")
	dbURL := os.Getenv("DB_URL")

	if port == "" {
		log.Fatal("PORT not found in environments")
	}
	if dbURL == "" {
		log.Fatal("DB_URL not found in environment")
	}

	db, err := sql.Open("postgres", dbURL)
	if err != nil {
		log.Fatal(err.Error())
	}

	dbQueries := database.New(db)

	cfg := apiConfig{DB: dbQueries}

	router := chi.NewRouter()
	router.Use(cors.Handler(cors.Options{}))
	router.Get("/err", getErr)

	v1Router := newV1Router(&cfg)
	router.Mount("/v1", v1Router)

	server := &http.Server{
		Handler: router,
		Addr:    ":" + port,
	}
	fmt.Println("Server listening on port:", port)
	server.ListenAndServe()
}
