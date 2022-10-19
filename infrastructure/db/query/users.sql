-- name: GetUsers :one
SELECT * FROM users WHERE id = $1 LIMIT 1;

-- name: ListUsers :many
SELECT * FROM users;

-- name: CreateUsers :one
INSERT INTO users (cpf, name, pass) VALUES ($1, $2, $3) RETURNING *;

-- name: UpdateUsers :exec
UPDATE users SET cpf = $2, name = $3, pass = $4 WHERE id = $1;

-- name: DeleteUsers :exec
DELETE FROM users WHERE id = $1;