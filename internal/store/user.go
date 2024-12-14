package store

import (
	"context"

	"github.com/stefanaki/shelter/internal/db"
)

type UserStore struct {
	queries *db.Queries
}

func (s UserStore) Create(ctx context.Context, user *db.CreateUserParams) (db.User, error) {
	newUser, err := s.queries.CreateUser(ctx, db.CreateUserParams{
		Username: user.Username,
		Password: user.Password,
		Email:    user.Email,
	})

	return newUser, err
}
