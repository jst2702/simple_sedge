-- name: CreateModel :one
INSERT INTO models (
  name, role, description
) VALUES (
  $1, $2, $3
)
RETURNING *;

-- name: GetModel :one
SELECT * FROM models
WHERE id = $1 LIMIT 1;

-- name: ListModels :many
SELECT * FROM models
ORDER by id
limit $1
OFFSET $2;

-- name: UpdateModel :one
UPDATE models
  set name = $2, 
  role = $3, 
  description = $4
WHERE id = $1
RETURNING *;

-- name: DeleteModel :exec
DELETE FROM models
WHERE id = $1;