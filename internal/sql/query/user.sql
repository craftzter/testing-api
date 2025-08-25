-- name: CreateUser :one
INSERT INTO users (
  username, email, password, created, updated
) VALUES (
  $1, $2, $3, $4, $5
) RETURNING *;

-- name: GetUser :one
SELECT * FROM users
WHERE id = $1
LIMIT 1;

-- name: ListUsers :many
SELECT * FROM users
ORDER BY id DESC
LIMIT 5;


-- name: GetUserByEmail :one
SELECT * FROM users
WHERE email = $1
LIMIT 1;
-- name: UpdateUser :one
UPDATE users
SET username = $2, email = $3, password = $4, updated = $5
WHERE id = $1
RETURNING *;
-- name: DeleteUser :exec
DELETE FROM users
WHERE id = $1;
