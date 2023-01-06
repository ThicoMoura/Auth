-- name: NewSession :one
INSERT INTO "session" ("id", "user", "token", "ip", "agent", "expires_at") VALUES ($1, $2, $3, $4, $5, $6) RETURNING *;

-- name: GetSession :one
SELECT * FROM "session" WHERE "id" = $1 LIMIT 1;

-- name: FindSession :many
SELECT * FROM "session" WHERE "user" = $1;

-- name: ListSession :many
SELECT * FROM "session" ORDER BY "created_at";

-- name: ListSessionPage :many
SELECT * FROM "session" ORDER BY "created_at" LIMIT $1 OFFSET $2;

-- name: DeleteSession :one
DELETE FROM "session" WHERE "id" = $1 RETURNING *;