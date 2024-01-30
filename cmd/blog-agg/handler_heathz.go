package main

import (
	"net/http"

	"github.com/PalmerTurley34/blog-aggregator/internal/responses"
)

func getReadiness(w http.ResponseWriter, r *http.Request) {
	type returnBody struct {
		Status string `json:"status"`
	}
	responses.WithJSON(w, http.StatusOK, returnBody{Status: "ok"})
}

func getErr(w http.ResponseWriter, r *http.Request) {
	responses.WithError(w, http.StatusInternalServerError, "Internal Server Error")
}
