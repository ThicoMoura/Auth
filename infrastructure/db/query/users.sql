-- name: GetUsers :one
SELECT * FROM users WHERE id = $1 LIMIT 1;

-- name: ListUsers :many
SELECT * FROM users ORDER BY id LIMIT $1 OFFSET $2;

-- name: CreateUsers :one
INSERT INTO users (cpf, name, pass) VALUES ($1, $2, $3) RETURNING *;

-- name: UpdateUsers :one
UPDATE users SET cpf = $2, name = $3, pass = $4, active = $5 WHERE id = $1 RETURNING *;

-- name: DeleteUsers :exec
DELETE FROM users WHERE id = $1;