-- name: StoreSession :exec
INSERT INTO sessions (id, user_id, expires_at)
VALUES ($1, $2, $3);

-- name: GetSession :one
SELECT *
FROM sessions
WHERE id = $1;
