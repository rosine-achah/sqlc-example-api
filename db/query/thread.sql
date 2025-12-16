-- -- name: CreateThread :one
-- INSERT INTO thread (title)
-- VALUES ($1)
-- RETURNING id, title, created_at;

-- -- name: GetThreadByID :one
-- SELECT id, title, created_at
-- FROM thread
-- WHERE id = $1;

-- -- name: ListThreads :many
-- SELECT id, title, created_at
-- FROM thread
-- ORDER BY created_at DESC;



-- name: CreateThread :one
INSERT INTO thread (title)
VALUES ($1)
RETURNING id, title, created_at;

-- name: GetThread :one
SELECT id, title, created_at
FROM thread
WHERE id = $1;

-- name: ListThreads :many
SELECT id, title, created_at
FROM thread
ORDER BY created_at DESC;
