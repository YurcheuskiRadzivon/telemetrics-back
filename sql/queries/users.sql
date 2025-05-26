-- name: CreateUser :exec
INSERT INTO users (user_id, username, phone_number)
VALUES ($1, $2, $3);

-- name: DeleteUser :exec
DELETE FROM users WHERE user_id = $1;

-- name: UpdateUser :exec
UPDATE users
SET username = $2, phone_number = $3
WHERE user_id = $1;

-- name: GetUser :one
SELECT * FROM users WHERE user_id = $1;
