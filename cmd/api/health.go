package main

import (
	"net/http"
)

func (a *api) healthcheckHandler(w http.ResponseWriter, r *http.Request) {
	data := map[string]string{
		"status":  "ok",
		"env":     a.config.env,
		"version": API_VERSION,
	}
	if err := writeJSON(w, http.StatusOK, data); err != nil {
		a.internalServerError(w, r, err)
	}
}
