-- name: CreatePost :one
INSERT INTO
    posts (
        id,
        created_at,
        updated_at,
        title,
        url,
        description,
        published_at,
        feed_id
    )
VALUES
    ($1, $2, $3, $4, $5, $6, $7, $8)
RETURNING
    *;

-- name: GetPostsForUser :many
SELECT
    posts.*
FROM
    posts
    INNER JOIN feeds_follow ON feeds_follow.feed_id = posts.feed_id
WHERE
    feeds_follow.user_id = $1
ORDER BY
    posts.updated_at DESC
LIMIT
    $2;