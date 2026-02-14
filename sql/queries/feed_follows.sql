-- name: CreateFeedFollow :exec
INSERT INTO feed_follows (id, created_at, updated_at, feed_id, user_id)
VALUES (
	$1, 
	$2, 
	$3, 
	$4, 
	$5
) RETURNING (
	SELECT * FROM feed_follows 
	JOIN users ON feed_follows.user_id = users.id 
	JOIN feeds ON feed_follows.feed_id = feeds.id 
);
