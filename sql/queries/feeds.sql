-- name: CreateFeed :one
INSERT INTO feeds (id, name, url, user_id, created_at, updated_at, last_fetched_at)
VALUES (
    $1,
    $2,
    $3,
    $4,
    $5,
    $6,
    $7
)
RETURNING *;

-- name: ListFeedsWithUsers :many
SELECT
    feeds.name AS feed_name,
    users.name AS user_name FROM feeds
INNER JOIN users
ON feeds.user_id = users.id;

-- name: GetFeedByURL :one
SELECT * from feeds
WHERE url = $1;

-- name: MarkFeedFetched :exec
UPDATE feeds
SET
    last_fetched_at = $2,
    updated_at = $2
WHERE id = $1;

-- name: GetNextFeedToFetch :one
SELECT id, url, name FROM feeds
ORDER BY feeds.last_fetched_at NULLS FIRST
LIMIT 1;
