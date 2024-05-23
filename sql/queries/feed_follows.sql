-- name: CreateFeedFollow :one
INSERT INTO feed_follows (id, user_id, feed_id) 
VALUES ($1, $2, $3) 
RETURNING *;

-- name: DeleteFeedFollowByID :exec
DELETE FROM feed_follows 
WHERE user_id = $1 AND feed_id = $2;

-- name: DeleteFeedFollowsByUserID :exec
DELETE FROM feed_follows 
WHERE user_id = $1;

-- name: GetFeedFollowsByUserID :many
SELECT * FROM feed_follows 
WHERE user_id = $1;


