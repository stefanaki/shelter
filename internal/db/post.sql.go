// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0
// source: post.sql

package db

import (
	"context"

	"github.com/jackc/pgx/v5/pgtype"
)

const createPost = `-- name: CreatePost :one
INSERT INTO posts (title, content, tags, user_id)
VALUES ($1, $2, $3, $4)
RETURNING id, title, content, tags, user_id, created_at, updated_at
`

type CreatePostParams struct {
	Title   string
	Content string
	Tags    []string
	UserID  pgtype.UUID
}

func (q *Queries) CreatePost(ctx context.Context, arg CreatePostParams) (Post, error) {
	row := q.db.QueryRow(ctx, createPost,
		arg.Title,
		arg.Content,
		arg.Tags,
		arg.UserID,
	)
	var i Post
	err := row.Scan(
		&i.ID,
		&i.Title,
		&i.Content,
		&i.Tags,
		&i.UserID,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const deletePost = `-- name: DeletePost :exec
DELETE FROM posts
WHERE id = $1
`

func (q *Queries) DeletePost(ctx context.Context, id pgtype.UUID) error {
	_, err := q.db.Exec(ctx, deletePost, id)
	return err
}

const getPostByID = `-- name: GetPostByID :one
SELECT id, title, content, tags, user_id, created_at, updated_at
FROM posts
WHERE id = $1
`

func (q *Queries) GetPostByID(ctx context.Context, id pgtype.UUID) (Post, error) {
	row := q.db.QueryRow(ctx, getPostByID, id)
	var i Post
	err := row.Scan(
		&i.ID,
		&i.Title,
		&i.Content,
		&i.Tags,
		&i.UserID,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const listPostsByUser = `-- name: ListPostsByUser :many
SELECT id, title, content, tags, user_id, created_at, updated_at
FROM posts
WHERE user_id = $1
ORDER BY created_at DESC
`

func (q *Queries) ListPostsByUser(ctx context.Context, userID pgtype.UUID) ([]Post, error) {
	rows, err := q.db.Query(ctx, listPostsByUser, userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Post
	for rows.Next() {
		var i Post
		if err := rows.Scan(
			&i.ID,
			&i.Title,
			&i.Content,
			&i.Tags,
			&i.UserID,
			&i.CreatedAt,
			&i.UpdatedAt,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const updatePost = `-- name: UpdatePost :one
UPDATE posts
SET
    title = COALESCE($2, title),
    content = COALESCE($3, content),
    tags = COALESCE($4, tags),
    user_id = COALESCE($5, user_id),
    updated_at = NOW()
WHERE id = $1
RETURNING id, title, content, tags, user_id, created_at, updated_at
`

type UpdatePostParams struct {
	ID      pgtype.UUID
	Title   string
	Content string
	Tags    []string
	UserID  pgtype.UUID
}

func (q *Queries) UpdatePost(ctx context.Context, arg UpdatePostParams) (Post, error) {
	row := q.db.QueryRow(ctx, updatePost,
		arg.ID,
		arg.Title,
		arg.Content,
		arg.Tags,
		arg.UserID,
	)
	var i Post
	err := row.Scan(
		&i.ID,
		&i.Title,
		&i.Content,
		&i.Tags,
		&i.UserID,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}
