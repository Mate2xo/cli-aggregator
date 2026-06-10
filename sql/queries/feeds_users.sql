-- name: CreateFeedFollow :one
WITH created_feed_follow AS (
  INSERT INTO feed_follows(user_id, feed_id, created_at, updated_at)
  VALUES ( $1, $2, NOW(), NOW() )
  RETURNING *
)
SELECT *,
  feeds.name AS feed_name,
  users.name AS user_name
FROM created_feed_follow
INNER JOIN feeds ON feeds.id = created_feed_follow.feed_id
INNER JOIN users ON users.id = created_feed_follow.user_id;

-- name: DeleteFeedFollow :exec
DELETE FROM feed_follows
WHERE user_id = $1 AND feed_id = $2;

-- name: GetFeedFollowsForUser :many
SELECT feed_follows.*, feeds.name as feed_name, users.name AS user_name
FROM feed_follows
INNER JOIN feeds ON feed_id = feeds.id
INNER JOIN users ON feed_follows.user_id = users.id
WHERE feed_follows.user_id = $1;

-- name: ResetFeedFollows :exec
DELETE FROM feed_follows;
