-- name: CreateFeedFollow :one
WITH insert_into_feed_follows AS (
	INSERT INTO feed_follows(id, created_at, updated_at, user_id, feed_id)
	VALUES ($1, $2, $3, $4, $5)
	RETURNING *
)

SELECT
	insert_into_feed_follows.*,
	feeds.name AS feed_name,
	users.name AS user_name
FROM
	insert_into_feed_follows
INNER JOIN feeds on insert_into_feed_follows.feed_id = feeds.id
INNER JOIN users on insert_into_feed_follows.user_id = users.id;

-- name: GetFeedFollowsForUser :many
WITH feed_follows_selected AS (
	SELECT * FROM feed_follows WHERE feed_follows.user_id = $1
)

SELECT
	feed_follows_selected.*,
	feeds.name AS feed_name,
	users.name AS user_name
FROM
	feed_follows_selected
INNER JOIN feeds on feed_follows_selected.feed_id = feeds.id
INNER JOIN users on feed_follows_selected.user_id = users.id;
