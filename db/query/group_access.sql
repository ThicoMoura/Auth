-- name: NewGroupAccess :one
INSERT INTO "group_access" ("group_id", "access_id") VALUES ($1, $2) RETURNING *;

-- name: GetGroupAccess :one
SELECT * FROM "group_access" WHERE "group_id" = $1 AND "access_id" = $2 LIMIT 1;

-- name: FindGroupAccess :many
SELECT * FROM "group_access" WHERE "group_id" = $1 OR "access_id" = $2;

-- name: FindGroupAccessPage :many
SELECT * FROM "group_access" WHERE "group_id" = $1 OR "access_id" = $2 LIMIT $3 OFFSET $4;

-- name: DeleteGroupAccess :one
DELETE FROM "group_access" WHERE "group_id" = $1 AND "access_id" = $2 RETURNING *;