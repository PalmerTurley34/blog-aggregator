package responses

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func WithJSON(w http.ResponseWriter, statusCode int, payload interface{}) {
	w.Header().Set("Content-Type", "application/json")
	data, err := json.Marshal(payload)
	if err != nil {
		fmt.Println("Error marshaling json")
		w.WriteHeader(500)
	}
	w.WriteHeader(statusCode)
	w.Write(data)
}

func WithError(w http.ResponseWriter, statusCode int, msg string) {
	type errorResponse struct {
		Error string `json:"error"`
	}
	WithJSON(w, statusCode, errorResponse{Error: msg})
}
