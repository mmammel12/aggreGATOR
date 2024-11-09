-- name: CreateFeed :one
INSERT INTO feeds (id, name, url, user_id)
VALUES (
    $1,
    $2,
    $3,
    $4
)
RETURNING *;

-- name: GetFeed :one
SELECT * FROM USERS
WHERE name = $1 LIMIT 1;

-- name: DeleteFeeds :exec
DELETE FROM feeds;

-- name: ListFeeds :many
SELECT * FROM feeds;

-- name: ListFeedsWithUsers :many
SELECT feeds.name, feeds.url, users.name AS User FROM feeds
INNER JOIN users
ON feeds.user_id = users.id;
