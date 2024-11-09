-- name: CreateFeed :one
INSERT INTO feeds (id, name, url, user_id)
VALUES (
    $1,
    $2,
    $3,
    $4
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
