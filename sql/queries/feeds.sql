-- name: GetFeeds :many
SELECT
    feeds.name,
    feeds.url,
    users.name as username
FROM
    feeds
    LEFT JOIN users ON feeds.user_id = users.id
ORDER BY
    users.name;

-- name: CreateFeed :one
INSERT INTO
    feeds (id, created_at, updated_at, name, url, user_id)
VALUES
    ($1, $2, $3, $4, $5, $6)
    RETURNING *;

-- name: CreateFeedFollow :one
WITH
    inserted_feed_follow AS (
        INSERT INTO
            feeds_follow (id, created_at, updated_at, user_id, feed_id)
        VALUES
            ($1, $2, $3, $4, $5)
            RETURNING *
    )
SELECT
    inserted_feed_follow.*,
    feeds.name as feed_name,
    users.name as user_name
FROM
    inserted_feed_follow
    LEFT JOIN users on inserted_feed_follow.user_id = users.id
    LEFT JOIN feeds on inserted_feed_follow.feed_id = feeds.id;

-- name: GetFeed :one
SELECT
    *
FROM
    feeds
WHERE
    url = $1;

-- name: GetFeedFollowsForUser :many
SELECT feeds_follow.*, users.name as user_name, feeds.name as feed_name FROM feeds_follow
LEFT JOIN users on feeds_follow.user_id = users.id
LEFT JOIN feeds on feeds_follow.feed_id = feeds.id
WHERE users.name = $1;