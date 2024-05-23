-- name: CreateFeed :one
INSERT INTO feeds (id, name, url, user_id)
VALUES ($1, $2, $3, $4)
RETURNING *;

-- name: GetAllFeeds :many
SELECT * FROM feeds
ORDER by updated_at DESC;

-- name: GetFeedById :one
SELECT * FROM feeds
WHERE id = $1;

-- name: UpdateFeedById :one
UPDATE feeds
SET name = $2, url = $3
WHERE id = $1
RETURNING *;

-- name: DeleteFeedById :exec
DELETE FROM feeds
WHERE id = $1;

-- name: GetNextFeedToFetch :one
SELECT *
FROM feeds
ORDER BY last_fetched_at IS NULL DESC, last_fetched_at ASC
LIMIT 1;

-- name: MarkFeedAsFetched :one
UPDATE feeds
SET last_fetched_at = CURRENT_TIMESTAMP,
    updated_at = CURRENT_TIMESTAMP
WHERE id = $1
RETURNING *;


