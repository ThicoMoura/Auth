-- name: NewUser :one
INSERT INTO "user" ("group", "email", "name", "pass") VALUES ($1, $2, $3, $4) RETURNING *;

-- name: GetUser :one
SELECT * FROM "user" WHERE "id" = $1 OR "email" = $2 LIMIT 1;

-- name: FindUser :many
SELECT * FROM  "user" WHERE "group" = $1 OR "name" LIKE $2 ORDER BY "name";

-- name: FindUserPage :many
SELECT * FROM  "user" WHERE "group" = $1 OR "name" LIKE $2 ORDER BY "name" LIMIT $3 OFFSET $4;

-- name: ListUser :many
SELECT * FROM "user" ORDER BY "name";

-- name: ListUserPage :many
SELECT * FROM "user" ORDER BY "name" LIMIT $1 OFFSET $2;

-- name: UpdateUser :one
UPDATE "user" SET "name" = COALESCE(NULLIF(@Name, ''), "name"), "pass" = COALESCE(NULLIF(@Pass, ''), "pass") WHERE "id" = $1 RETURNING *;

-- name: UpdateUserActive :one
UPDATE "user" SET "active" = $2 WHERE "id" = $1 RETURNING *;

-- name: DeleteUser :one
DELETE FROM "user" WHERE "id" = $1 RETURNING *;