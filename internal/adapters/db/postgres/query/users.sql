-- name: StoreUser :exec
INSERT INTO users (first_name, last_name, phone, email, password)
VALUES ($1, $2, $3, $4, $5)
RETURNING id;

-- name: GetUserByID :one
SELECT *
FROM users
WHERE id = $1;

-- name: GetUserByEmail :one
SELECT *
FROM users
WHERE email = $1;

-- name: SearchUserByEmail :many
SELECT *
FROM users
WHERE email LIKE $1;

-- name: GetAllUsers :many
SELECT id,
       first_name,
       last_name
FROM users;
