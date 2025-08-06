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
