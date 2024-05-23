-- name: CreatePost :one
INSERT INTO posts (title, url, description, published_at, feed_id)
VALUES ($1, $2, $3, $4, $5)
RETURNING *;

-- name: GetPostsByUserID :many
SELECT p.*
FROM posts p
JOIN feed_follows ff ON p.feed_id = ff.feed_id
WHERE ff.user_id = $1
ORDER BY p.updated_at DESC
LIMIT $2
OFFSET $3; 

