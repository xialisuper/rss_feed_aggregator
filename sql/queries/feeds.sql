-- name: CreateFeed :one
INSERT INTO feeds (id, name, url, user_id)
VALUES ($1, $2, $3, $4)
RETURNING *;

-- name: GetAllFeeds :many
SELECT * FROM feeds
ORDER by updated_at DESC;
