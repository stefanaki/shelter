-- name: ListCommentsByPostID :many
SELECT c.id, c.user_id, c.content, c.created_at
FROM comments c
LEFT JOIN posts p on c.post_id = p.id
WHERE p.id = $1;