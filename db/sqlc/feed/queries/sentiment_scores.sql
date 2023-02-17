-- name: CreateScore :one
INSERT INTO sentiment_scores (
  model_id, document_guid, sentiment, confidence
) VALUES (
  $1, $2, $3, $4
)
RETURNING *;

-- name: GetScore :one
SELECT * FROM sentiment_scores
WHERE id = $1 LIMIT 1;

-- name: ListScores :many
SELECT * FROM sentiment_scores
ORDER by id
limit $1
OFFSET $2;

-- name: UpdateScore :one
UPDATE sentiment_scores
  set sentiment = $2,
  confidence = $3
WHERE id = $1
RETURNING *;

-- name: DeleteScore :exec
DELETE FROM sentiment_scores
WHERE id = $1;