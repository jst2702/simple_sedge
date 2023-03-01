-- name: CreateDocument :one
INSERT INTO documents (
  guid,
  url, 
  site, 
  site_full, 
  site_section, 
  headline, 
  title, 
  body, 
  ticker, 
  tickers, 
  published_at, 
  language
) VALUES (
  $1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12
)
RETURNING *;

-- name: GetDocument :one
SELECT * FROM documents
WHERE guid = $1 LIMIT 1;

-- name: ListDocuments :many
SELECT * FROM documents
ORDER by guid
limit $1
OFFSET $2;

-- name: UpdateDocument :one
UPDATE documents
  set ticker = $2, 
  tickers = $3
WHERE guid = $1
RETURNING *;

-- name: DeleteDocuemnt :exec
DELETE FROM documents
WHERE guid = $1;