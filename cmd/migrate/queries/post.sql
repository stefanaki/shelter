-- name: CreatePost :one
INSERT INTO posts (title, content, tags, user_id)
VALUES ($1, $2, $3, $4)
RETURNING *;

-- name: GetPostByID :one
SELECT id, title, content, tags, user_id, created_at, updated_at
FROM posts
WHERE id = $1;

-- name: ListPostsByUser :many
SELECT id, title, content, tags, user_id, created_at, updated_at
FROM posts
WHERE user_id = $1
ORDER BY created_at DESC;