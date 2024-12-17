package store

import (
	"context"
	"database/sql"
	"errors"
	"fmt"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgtype"

	"github.com/stefanaki/shelter/internal/db"
	"github.com/stefanaki/shelter/internal/dto"
)

type PostStore struct {
	queries *db.Queries
}

func (s PostStore) Create(ctx context.Context, post dto.CreatePostPayload) (*db.Post, error) {
	parsedUUID, err := uuid.Parse(post.UserID)
	if err != nil {
		return nil, fmt.Errorf("invalid user_id: %w", err)
	}

	newPost, err := s.queries.CreatePost(ctx, db.CreatePostParams{
		Title:   post.Title,
		Content: post.Content,
		Tags:    post.Tags,
		UserID:  pgtype.UUID{Bytes: parsedUUID, Valid: true},
	})

	return &newPost, err
}

func (s PostStore) Retrieve(ctx context.Context, id string) (*db.Post, error) {
	postId, err := uuid.Parse(id)
	if err != nil {
		return nil, err
	}

	post, err := s.queries.GetPostByID(ctx, pgtype.UUID{Bytes: postId, Valid: true})
	if err != nil {
		switch {
		case errors.Is(err, sql.ErrNoRows):
			return nil, ErrorNotFound
		default:
			return nil, err
		}
	}

	return &post, nil
}

func (s *PostStore) Delete(ctx context.Context, id string) error {
	postID, err := StringToPostgresUUIDColumn(id)
	if err != nil {
		return ErrorInvalidInput
	}

	err = s.queries.DeletePost(ctx, postID)
	if err != nil {
		return err
	}

	return nil
}

func (s *PostStore) Update(ctx context.Context, id string, payload dto.UpdatePostPayload) (*db.Post, error) {
	postID, _ := StringToPostgresUUIDColumn(id)
	userID, _ := StringToPostgresUUIDColumn(payload.UserID)

	updatedPost, err := s.queries.UpdatePost(ctx, db.UpdatePostParams{
		ID:      postID,
		Title:   payload.Title,
		Content: payload.Content,
		Tags:    payload.Tags,
		UserID:  userID,
	})

	return &updatedPost, err
}
