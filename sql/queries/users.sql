-- name: CreateUser :one
INSERT INTO users (id, name)
VALUES ($1, $2)
RETURNING *;

-- name: GetUserByApiKey :one
SELECT *
FROM users
WHERE api_key = $1;

-- name: UpdateUser :one
UPDATE users 
    SET name = $2 
WHERE api_key = $1
RETURNING *;






