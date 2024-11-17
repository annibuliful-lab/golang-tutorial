package utils

import (
	"encoding/json"
	"errors"
	"io"
	"net/http"
)

func RespondWithJSON(w http.ResponseWriter, statusCode int, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	json.NewEncoder(w).Encode(data)
}

func ProcessRequestBody(r *http.Request, target interface{}) error {
	// Read the request body
	body, err := io.ReadAll(r.Body)
	if err != nil {
		return errors.New("failed to read request body")
	}
	defer r.Body.Close() // Ensure the body is closed

	// Decode the JSON body into the target struct
	if err := json.Unmarshal(body, target); err != nil {
		return errors.New("invalid JSON format")
	}

	return nil
}