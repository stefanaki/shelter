package main

import (
	"log"
	"net/http"
	"time"
)

type api struct {
	config config
}

type config struct {
	address string
}

func (a *api) mount() *http.ServeMux {
	mux := http.NewServeMux()

	mux.HandleFunc("GET /v1/health", healthcheckHandler)

	return mux
}

func (a *api) serve(mux *http.ServeMux) error {
	srv := &http.Server{
		Addr:         a.config.address,
		Handler:      mux,
		WriteTimeout: time.Second * 30,
		ReadTimeout:  time.Second * 10,
		IdleTimeout:  time.Minute,
	}

	log.Printf("Server listening on %v...\n", srv.Addr)

	return srv.ListenAndServe()
}
