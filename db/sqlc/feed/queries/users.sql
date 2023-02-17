-- name: CreateUser :one
INSERT INTO users (
  email, hashed_password
) VALUES (
  $1, $2
)
RETURNING *;

-- name: GetUser :one
SELECT * FROM users
WHERE email = $1 LIMIT 1;

-- name: ListUsers :many
SELECT * FROM users
ORDER by email
limit $1
OFFSET $2;

-- name: UpdateUser :one
UPDATE users set hashed_password = $2
WHERE email = $1
RETURNING *;

-- name: DeleteUser :exec
DELETE FROM users
WHERE email = $1;