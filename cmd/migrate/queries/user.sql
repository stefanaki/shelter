-- name: CreateUser :one
INSERT INTO users (username, password, email)
VALUES ($1, $2, $3)
RETURNING *;

-- name: GetUserByID :one
SELECT id, username, password, email, created_at
FROM users
WHERE id = $1;