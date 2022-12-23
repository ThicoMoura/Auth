// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.16.0
// source: group.sql

package db

import (
	"context"

	"github.com/google/uuid"
)

const DeleteGroup = `-- name: DeleteGroup :one
DELETE FROM "group" WHERE "id" = $1 RETURNING id, name, active
`

func (q *Queries) DeleteGroup(ctx context.Context, db DBTX, id uuid.UUID) (*Group, error) {
	row := db.QueryRow(ctx, DeleteGroup, id)
	var i Group
	err := row.Scan(&i.ID, &i.Name, &i.Active)
	return &i, err
}

const GetGroup = `-- name: GetGroup :one
SELECT id, name, active FROM "group" WHERE "id" = $1 LIMIT 1
`

func (q *Queries) GetGroup(ctx context.Context, db DBTX, id uuid.UUID) (*Group, error) {
	row := db.QueryRow(ctx, GetGroup, id)
	var i Group
	err := row.Scan(&i.ID, &i.Name, &i.Active)
	return &i, err
}

const ListGroup = `-- name: ListGroup :many
SELECT id, name, active FROM "group" ORDER BY "name"
`

func (q *Queries) ListGroup(ctx context.Context, db DBTX) ([]*Group, error) {
	rows, err := db.Query(ctx, ListGroup)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []*Group{}
	for rows.Next() {
		var i Group
		if err := rows.Scan(&i.ID, &i.Name, &i.Active); err != nil {
			return nil, err
		}
		items = append(items, &i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const ListGroupPage = `-- name: ListGroupPage :many
SELECT id, name, active FROM "group" ORDER BY "name" LIMIT $1 OFFSET $2
`

type ListGroupPageParams struct {
	Limit  int32 `db:"limit" json:"limit"`
	Offset int32 `db:"offset" json:"offset"`
}

func (q *Queries) ListGroupPage(ctx context.Context, db DBTX, arg *ListGroupPageParams) ([]*Group, error) {
	rows, err := db.Query(ctx, ListGroupPage, arg.Limit, arg.Offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []*Group{}
	for rows.Next() {
		var i Group
		if err := rows.Scan(&i.ID, &i.Name, &i.Active); err != nil {
			return nil, err
		}
		items = append(items, &i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const NewGroup = `-- name: NewGroup :one
INSERT INTO "group" ("name") VALUES ($1) RETURNING id, name, active
`

func (q *Queries) NewGroup(ctx context.Context, db DBTX, name string) (*Group, error) {
	row := db.QueryRow(ctx, NewGroup, name)
	var i Group
	err := row.Scan(&i.ID, &i.Name, &i.Active)
	return &i, err
}

const UpdateGroup = `-- name: UpdateGroup :one
UPDATE "group" SET "name" = $2, "active" = $3 WHERE "id" = $1 RETURNING id, name, active
`

type UpdateGroupParams struct {
	ID     uuid.UUID `db:"id" json:"id"`
	Name   string    `db:"name" json:"name"`
	Active bool      `db:"active" json:"active"`
}

func (q *Queries) UpdateGroup(ctx context.Context, db DBTX, arg *UpdateGroupParams) (*Group, error) {
	row := db.QueryRow(ctx, UpdateGroup, arg.ID, arg.Name, arg.Active)
	var i Group
	err := row.Scan(&i.ID, &i.Name, &i.Active)
	return &i, err
}
