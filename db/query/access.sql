-- name: NewAccess :one
INSERT INTO "access" ("system", "table", "type") VALUES ($1, $2, $3) RETURNING *;

-- name: GetAccess :one
SELECT * FROM "access" WHERE "id" = $1 LIMIT 1;

-- name: FindAccess :many
SELECT * FROM  "access" WHERE "system" = $1 OR "table" LIKE $2 ORDER BY "table";

-- name: FindAccessPage :many
SELECT * FROM  "access" WHERE "system" = $1 OR "table" LIKE $2 ORDER BY "table" LIMIT $3 OFFSET $4;

-- name: ListAccess :many
SELECT * FROM "access" ORDER BY "table";

-- name: ListAccessPage :many
SELECT * FROM "access" ORDER BY "table" LIMIT $1 OFFSET $2;

-- name: UpdateAccess :one
UPDATE "access" SET "type" = $2 WHERE "id" = $1 RETURNING *;

-- name: DeleteAccess :one
DELETE FROM "access" WHERE "id" = $1 RETURNING *;