-- name: CreateMessage :one
INSERT INTO message (thread, sender, content)
VALUES ($1, $2, $3)
RETURNING id, thread, sender, content, created_at, updated_at;

-- name: GetMessageByID :one
SELECT id, thread, sender, content, created_at, updated_at
FROM message
WHERE id = $1;

-- name: GetMessagesByThreadPaginated :many
SELECT id, thread, sender, content, created_at
FROM message
WHERE thread = $1
ORDER BY created_at ASC
LIMIT $2::int
OFFSET $3::int;

-- name: UpdateMessageContent :one
UPDATE message
SET content = $2, updated_at = now()
WHERE id = $1
RETURNING id, thread, sender, content, created_at, updated_at;

-- name: DeleteMessageByID :exec
DELETE FROM message
WHERE id = $1;
