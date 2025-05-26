-- name: CreateViewOptions :exec
INSERT INTO view_options (
    user_id, channel_count, tittle, about, channel_id,
    channel_date, participants_count, photo, message_count,
    message_id, views, post_date, reactions_count, reactions
) VALUES (
    $1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14
);

-- name: UpdateViewOptions :exec
UPDATE view_options
SET
    channel_count = $2,
    tittle = $3,
    about = $4,
    channel_id = $5,
    channel_date = $6,
    participants_count = $7,
    photo = $8,
    message_count = $9,
    message_id = $10,
    views = $11,
    post_date = $12,
    reactions_count = $13,
    reactions = $14
WHERE user_id = $1;

-- name: GetViewOptions :one
SELECT * FROM view_options WHERE user_id = $1;

-- name: DeleteViewOptions :exec
DELETE FROM view_options WHERE user_id = $1;