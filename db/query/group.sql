-- name: NewGroup :one
INSERT INTO "group" ("name") VALUES ($1) RETURNING *;

-- name: GetGroup :one
SELECT * FROM "group" WHERE "id" = $1 LIMIT 1;

-- name: ListGroup :many
SELECT * FROM "group" ORDER BY "name";

-- name: ListGroupPage :many
SELECT * FROM "group" ORDER BY "name" LIMIT $1 OFFSET $2;

-- name: UpdateGroup :one
UPDATE "group" SET "name" = $2, "active" = $3 WHERE "id" = $1 RETURNING *;

-- name: DeleteGroup :one
DELETE FROM "group" WHERE "id" = $1 RETURNING *;