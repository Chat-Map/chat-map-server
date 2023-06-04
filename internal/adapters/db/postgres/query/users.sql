-- name: StoreUser :exec
INSERT INTO users (first_name, last_name, email, password)
VALUES ($1, $2, $3, $4);
--
-- name: GetUser :one
SELECT *
FROM users
WHERE id = $1;
--
-- name: SearchUserByEmail :many
SELECT *
FROM users
WHERE email LIKE $1;
