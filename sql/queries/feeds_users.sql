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

-- name: GetFeedFollowsForUser :many
WITH with_user AS (
  SELECT * FROM users WHERE users.name = $1
)
SELECT feed_follows.*, feeds.name as feed_name, with_user.name AS user_name
FROM feed_follows
INNER JOIN feeds ON feed_id = feeds.id
INNER JOIN with_user ON feed_follows.user_id = with_user.id
WHERE feed_follows.user_id = with_user.id;

-- name: ResetFeedFollows :exec
DELETE FROM feed_follows;
