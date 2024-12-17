package store

import (
	"context"
	"errors"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/stefanaki/shelter/internal/db"
	"github.com/stefanaki/shelter/internal/dto"
)

var (
	ErrorInvalidInput = errors.New("invalid input")
	ErrorNotFound     = errors.New("resource not found")
)

type Store struct {
	Posts interface {
		Create(context.Context, dto.CreatePostPayload) (*db.Post, error)
		Retrieve(context.Context, string) (*db.Post, error)
		Update(context.Context, string, dto.UpdatePostPayload) (*db.Post, error)
		Delete(context.Context, string) error
	}
	Users interface {
		Create(context.Context, *db.CreateUserParams) (db.User, error)
	}
	Comments interface {
		ListByPostID(context.Context, string) (*[]db.Comment, error)
	}
}

func NewStore(pool *pgxpool.Pool) Store {
	db := db.New(pool)

	return Store{
		Posts:    &PostStore{db},
		Users:    &UserStore{db},
		Comments: &CommentStore{db},
	}
}
