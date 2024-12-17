package main

import (
	"encoding/json"
	"net/http"

	validator "github.com/go-playground/validator/v10"
)

const MAX_BYTES = int64(1_048_578)

var Validate *validator.Validate

func init() {
	Validate = validator.New(validator.WithRequiredStructEnabled())
}

func writeJSON(w http.ResponseWriter, status int, data any) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	return json.NewEncoder(w).Encode(data)
}

func readJSON(w http.ResponseWriter, r *http.Request, data any) error {
	r.Body = http.MaxBytesReader(w, r.Body, MAX_BYTES)

	decoder := json.NewDecoder(r.Body)
	decoder.DisallowUnknownFields()

	return decoder.Decode(data)
}

func writeError(w http.ResponseWriter, status int, message string) error {
	type res struct {
		Error string `json:"error"`
	}

	return writeJSON(w, status, res{Error: message})
}
