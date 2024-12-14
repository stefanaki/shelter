package store

import (
	"context"

	"github.com/stefanaki/shelter/internal/db"
)

type PostStore struct {
	queries *db.Queries
}

func (s PostStore) Create(ctx context.Context, post *db.CreatePostParams) (db.Post, error) {
	newPost, err := s.queries.CreatePost(ctx, db.CreatePostParams{
		Title:   post.Title,
		Content: post.Content,
		Tags:    post.Tags,
		UserID:  post.UserID,
	})

	return newPost, err
}
