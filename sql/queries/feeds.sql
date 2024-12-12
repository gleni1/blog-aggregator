-- name: CreateFeed :one
INSERT INTO feeds (id, created_at, updated_at, name, url, user_id)
VALUES (
  $1,
  $2,
  $3,
  $4,
  $5,
  $6
)
RETURNING *;

-- name: ListFeeds :many
SELECT feeds.name, url, users.name
FROM users 
INNER JOIN feeds 
ON users.id = feeds.user_id;

-- name: GetFeedByURL :one 
SELECT feeds.ID, feeds.Name 
FROM feeds 
WHERE feeds.URL = $1;

-- name: MarkFeedFetched :one
UPDATE feeds
SET last_fetched_at = NOW(), 
updated_at = NOW() 
WHERE feeds.id = $1
RETURNING *;

-- name: GetNextFeedToFetch :one 
SELECT * 
FROM feeds
ORDER BY feeds.last_fetched_at NULLS FIRST
LIMIT 1;
