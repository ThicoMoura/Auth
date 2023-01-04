-- name: NewSystem :one
INSERT INTO "system" ("name", "tables") VALUES ($1, $2) RETURNING *;

-- name: GetSystem :one
SELECT * FROM "system" WHERE "id" = $1 LIMIT 1;

-- name: FindSystem :many
SELECT * FROM "system" WHERE "name" LIKE $1 ORDER BY "name";

-- name: FindSystemPage :many
SELECT * FROM "system" WHERE "name" LIKE $1 ORDER BY "name" LIMIT $2 OFFSET $3;

-- name: ListSystem :many
SELECT * FROM "system" ORDER BY "name";

-- name: ListSystemPage :many
SELECT * FROM "system" ORDER BY "name" LIMIT $1 OFFSET $2;

-- name: UpdateSystem :one
UPDATE "system" SET "name" = COALESCE(NULLIF(@Name, ''), "name"), "tables" = COALESCE($2, "tables") WHERE "id" = $1 RETURNING *;

-- name: UpdateActiveSystem :one
UPDATE "system" SET "active" = $2 WHERE "id" = $1 RETURNING *;

-- name: DeleteSystem :one
DELETE FROM "system" WHERE "id" = $1 RETURNING *;