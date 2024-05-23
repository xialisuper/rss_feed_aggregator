-- name: CreatePost :exec
INSERT INTO posts (title, url, description, published_at, feed_id)
VALUES ($1, $2, $3, $4, $5)
ON CONFLICT (url) DO UPDATE
SET title = EXCLUDED.title,
    description = EXCLUDED.description,
    published_at = EXCLUDED.published_at,
    feed_id = EXCLUDED.feed_id;


-- name: GetPostsByUserID :many
SELECT p.*
FROM posts p
JOIN feed_follows ff ON p.feed_id = ff.feed_id
WHERE ff.user_id = $1
ORDER BY p.updated_at DESC
LIMIT $2
OFFSET $3; 

