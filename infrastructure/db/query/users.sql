-- name: GetUser :one
SELECT * FROM users WHERE id = $1 LIMIT 1;

-- name: GetUserByCPF :one
SELECT * FROM users WHERE cpf = $1 LIMIT 1;

-- name: ListUser :many
SELECT * FROM users ORDER BY id LIMIT $1 OFFSET $2;

-- name: CreateUser :one
INSERT INTO users (cpf, name, pass) VALUES ($1, $2, $3) RETURNING *;

-- name: UpdateUserName :one
UPDATE users SET name = $2 WHERE id = $1 RETURNING *;

-- name: UpdateUserPass :one
UPDATE users SET pass = $2 WHERE id = $1 RETURNING *;

-- name: UpdateUserActive :one
UPDATE users SET active = $2 WHERE id = $1 RETURNING *;

-- name: DeleteUser :exec
DELETE FROM users WHERE id = $1;