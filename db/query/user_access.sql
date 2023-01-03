-- name: NewUserAccess :one
INSERT INTO "user_access" ("user_id", "access_id") VALUES ($1, $2) RETURNING *;

-- name: GetUserAccess :one
SELECT * FROM "user_access" WHERE "user_id" = $1 AND "access_id" = $2 LIMIT 1;

-- name: FindUserAccess :many
SELECT * FROM "user_access" WHERE "user_id" = $1 OR "access_id" = $2;

-- name: FindUserAccessPage :many
SELECT * FROM "user_access" WHERE "user_id" = $1 OR "access_id" = $2 LIMIT $3 OFFSET $4;

-- name: DeleteUserAccess :one
DELETE FROM "user_access" WHERE "user_id" = $1 AND "access_id" = $2 RETURNING *;