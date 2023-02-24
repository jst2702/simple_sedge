-- name: CreateUser :one
INSERT INTO users (
  email, username, hashed_password
) VALUES (
  $1, $2, $3
)
RETURNING *;

-- name: GetUser :one
SELECT * FROM users
WHERE username = $1 LIMIT 1;

-- name: ListUsers :many
SELECT * FROM users
ORDER by username
limit $1
OFFSET $2;

-- name: UpdatePassword :one
UPDATE users set hashed_password = $2
WHERE username = $1
RETURNING *;

-- name: DeleteUser :exec
DELETE FROM users
WHERE email = $1;