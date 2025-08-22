-- name: GetUser :one
SELECT * FROM users
WHERE id = $1 LIMIT 1;

-- name: ListUsers :many
SELECT * FROM users
ORDER BY username;

-- name: CreateUser :one
INSERT INTO users (
  username, password_hash
) VALUES (
  $1, $2
)
RETURNING *;

-- name: UpdateUserName :exec
UPDATE users
  set username = $2
WHERE id = $1
RETURNING *;

-- name: UpdateUserPassword :exec
UPDATE users
  set password_hash = $2
WHERE id = $1
RETURNING *;

-- name: DeleteUser :exec
DELETE FROM users
WHERE id = $1;

-- name: GetNote :one
SELECT * FROM notes
WHERE id = $1 AND user_id = $2
LIMIT 1;

-- name: ListNotes :many
SELECT * FROM notes
WHERE user_id = $1
ORDER BY created_at DESC;

-- name: CreateNote :one
INSERT INTO notes (
  user_id, title, content
) VALUES (
  $1, $2, $3
)
RETURNING *;

-- name: UpdateNote :exec
UPDATE notes
  set title = $2,
  content = $3
WHERE id = $1
RETURNING *;

-- name: DeleteNote :exec
DELETE FROM notes
WHERE id = $1;

