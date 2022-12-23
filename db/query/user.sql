-- name: NewUser :one
INSERT INTO "user" ("group", "cpf", "name", "pass") VALUES ($1, $2, $3, $4) RETURNING *;

-- name: GetUser :one
SELECT * FROM "user" WHERE "id" = $1 LIMIT 1;

-- name: GetUserByCPF :one
SELECT * FROM "user" WHERE "cpf" = $1 LIMIT 1;

-- name: GetUserByGroup :many
SELECT * FROM  "user" WHERE "group" = $1 ORDER BY "name";

-- name: GetUserByGroupPage :many
SELECT * FROM  "user" WHERE "group" = $1 ORDER BY "name" LIMIT $2 OFFSET $3;

-- name: ListUser :many
SELECT * FROM "user" ORDER BY "name";

-- name: ListUserPage :many
SELECT * FROM "user" ORDER BY "name" LIMIT $1 OFFSET $2;

-- name: UpdateUser :one
UPDATE "user" SET "name" = $2, "active" = $3 WHERE "id" = $1 RETURNING *;

-- name: UpdateUserPass :one
UPDATE "user" SET "pass" = $2 WHERE "id" = $1 RETURNING *;

-- name: DeleteUser :one
DELETE FROM "user" WHERE "id" = $1 RETURNING *;