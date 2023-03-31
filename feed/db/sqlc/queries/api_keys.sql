-- name: CreateApiKey :one
INSERT INTO api_keys (
    api_key
) VALUES (
    $1
) RETURNING *;

-- name: GetApiKey :one
SELECT * FROM api_keys
WHERE api_key = $1 LIMIT 1;

-- name: DisableApiKey :one
UPDATE api_keys
SET
    active = FALSE
WHERE api_key = @api_key
RETURNING *;