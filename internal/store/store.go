package store

import (
	"context"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/stefanaki/shelter/internal/db"
)

type Store struct {
	Posts interface {
		Create(context.Context, *db.CreatePostParams) (db.Post, error)
	}
	Users interface {
		Create(context.Context, *db.CreateUserParams) (db.User, error)
	}
}

func NewStore(pool *pgxpool.Pool) Store {
	db := db.New(pool)

	return Store{
		Posts: &PostStore{db},
		Users: &UserStore{db},
	}
}
