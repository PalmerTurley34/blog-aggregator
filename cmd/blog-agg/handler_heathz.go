package main

import (
	"net/http"
)

func getReadiness(w http.ResponseWriter, r *http.Request) {
	type returnBody struct {
		Status string `json:"status"`
	}
	respondWithJSON(w, http.StatusOK, returnBody{Status: "ok"})
}

func getErr(w http.ResponseWriter, r *http.Request) {
	respondWithError(w, http.StatusInternalServerError, "Internal Server Error")
}
