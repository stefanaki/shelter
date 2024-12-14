package main

import (
	"context"
	"log"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"

	"github.com/stefanaki/shelter/internal/env"
	"github.com/stefanaki/shelter/internal/store"
)

func main() {
	cfg := config{
		address: env.GetString("ADDR", ":8080"),
		db: dbConfig{
			address:            env.GetString("DB_URI", "postgres://postgres:postgres@localhost/shelter?sslmode=disable"),
			maxOpenConnections: env.GetInt("DB_MAX_OPEN_CONNS", 30),
			maxIdleTime:        env.GetDuration("DB_MAX_IDLE_CONNS", 15*time.Minute),
		},
	}

	poolConfig, err := pgxpool.ParseConfig(cfg.db.address)
	if err != nil {
		log.Fatalf("failed to parse db config: %v", err)
	}

	poolConfig.MaxConns = int32(cfg.db.maxOpenConnections)
	poolConfig.MaxConnIdleTime = cfg.db.maxIdleTime

	pool, err := pgxpool.NewWithConfig(context.Background(), poolConfig)
	if err != nil {
		log.Fatalf("error creating db connection: %v", err)
	}

	defer pool.Close()
	log.Printf("Database connection pool established")

	store := store.NewStore(pool)
	app := &api{
		config: cfg,
		store:  store,
	}

	mux := app.mount()

	if err := app.serve(mux); err != nil {
		log.Fatalf("error serving api: %v", err)
	}
}
