-- name: CreateUser :one
INSERT INTO users (username, password, email)
VALUES ($1, $2, $3)
RETURNING id, username, password, email, created_at;

-- name: CreatePost :one
INSERT INTO posts (title, content, tags, user_id)
VALUES ($1, $2, $3, $4)
RETURNING id, title, content, tags, user_id, created_at, updated_at;

-- name: GetUserByID :one
SELECT id, username, password, email, created_at
FROM users
WHERE id = $1;

-- name: GetPostByID :one
SELECT id, title, content, tags, user_id, created_at, updated_at
FROM posts
WHERE id = $1;

-- name: ListPostsByUser :many
SELECT id, title, content, tags, user_id, created_at, updated_at
FROM posts
WHERE user_id = $1
ORDER BY created_at DESC;
