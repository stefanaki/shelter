package main

import (
	"log"
	"net/http"
)

func (a *api) internalServerError(w http.ResponseWriter, r *http.Request, err error) {
	log.Printf("[%s] %s Internal error: %v\n", r.Method, r.URL.Path, err)
	writeError(w, http.StatusInternalServerError, "Internal Server Error")
}

func (a *api) badRequest(w http.ResponseWriter, r *http.Request, err error) {
	log.Printf("[%s] %s Bad request: %v\n", r.Method, r.URL.Path, err)
	writeError(w, http.StatusBadRequest, err.Error())
}

func (a *api) notFound(w http.ResponseWriter, r *http.Request, err error) {
	log.Printf("[%s] %s Not found: %v\n", r.Method, r.URL.Path, err)
	writeError(w, http.StatusNotFound, "Not Found")
}
