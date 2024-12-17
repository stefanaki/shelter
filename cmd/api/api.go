package main

import (
	"log"
	"net/http"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/stefanaki/shelter/internal/store"
)

type api struct {
	config config
	store  store.Store
}

type config struct {
	env     string
	address string
	db      dbConfig
}

type dbConfig struct {
	address            string
	maxOpenConnections int
	maxIdleConnections int
	maxIdleTime        time.Duration
}

func (a *api) mount() http.Handler {
	r := chi.NewRouter()

	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Use(middleware.Timeout(60 * time.Second))

	r.Route("/v1", func(r chi.Router) {
		r.Get("/health", a.healthcheckHandler)

		r.Route("/users", func(r chi.Router) {
			r.Post("/", a.createUserHandler)
		})

		r.Route("/posts", func(r chi.Router) {
			r.Post("/", a.createPostHandler)

			r.Route("/{id}", func(r chi.Router) {
				r.Use(a.withPostContext)
				r.Get("/", a.retrievePostHandler)
				r.Patch("/", a.updatePostHandler)
				r.Delete("/", a.deletePostHandler)
				r.Get("/comments", a.listPostCommentsHandler)
			})

		})
	})

	return r
}

func (a *api) serve(mux http.Handler) error {
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
