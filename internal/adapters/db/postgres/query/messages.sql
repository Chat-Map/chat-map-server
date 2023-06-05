-- name: StoreMessage :one
INSERT INTO messages(chat_id, sender_id, content)
VALUES ($1, $2, $3)
RETURNING id;
