package store

import (
	"context"
	"database/sql"
	"errors"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgtype"
	"github.com/stefanaki/shelter/internal/db"
)

type CommentStore struct {
	queries *db.Queries
}

func (s CommentStore) ListByPostID(ctx context.Context, postID string) (*[]db.Comment, error) {
	id := pgtype.UUID{Bytes: uuid.MustParse(postID), Valid: true}

	comments, err := s.queries.ListCommentsByPostID(ctx, id)

	switch {
	case errors.Is(err, sql.ErrNoRows):
		return nil, ErrorNotFound
	}

	var res []db.Comment
	for _, c := range comments {
		res = append(res, db.Comment{
			ID:        c.ID,
			PostID:    id,
			UserID:    c.UserID,
			Content:   c.Content,
			CreatedAt: c.CreatedAt,
		})
	}

	return &res, nil
}
