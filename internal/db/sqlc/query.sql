-- name: GetUser :one
SELECT * FROM users where users.id = ?;

-- name: CreateUser :one
INSERT INTO users (
  username,
  'image',
  'password'
) VALUES (
  ?,
  ?,
  ?
)
RETURNING *;

