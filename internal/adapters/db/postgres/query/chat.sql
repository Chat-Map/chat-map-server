-- name: CreateChat :one
INSERT INTO chat(chat_type)
VALUES ($1)
RETURNING id;

-- name: AddChatMember :exec
INSERT INTO chat_users(chat_id, user_id)
VALUES ($1, $2);

-- name: GetChatUserRow :one
SELECT *
FROM chat_users cu
WHERE cu.chat_id = $1
  AND cu.user_id = $2
LIMIT 1;

-- name: GetChatMessages :many
SELECT *
FROM "messages" m
WHERE m.chat_id = $1
ORDER BY m.created_at DESC;

-- name: GetChatMembers :many
SELECT cu.user_id
FROM "chat_users" cu
WHERE cu.chat_id = $1;

-- name: GetUserChatMetadata :many
SELECT ch.id,
       u.first_name,
       u.last_name,
       COALESCE((SELECT m.content FROM "messages" m WHERE m.id = ch.id ORDER BY m.created_at DESC LIMIT 1),
                '')::varchar AS last_message
FROM "users" u
       JOIN "chat_users" cu ON u.ID = cu.user_id
       JOIN "chat" ch ON ch.id = cu.chat_id
WHERE u.id = $1;
