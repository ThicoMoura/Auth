// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.16.0
// source: user_access.sql

package db

import (
	"context"

	"github.com/google/uuid"
)

const DeleteUserAccess = `-- name: DeleteUserAccess :one
DELETE FROM "user_access" WHERE "user_id" = $1 AND "access_id" = $2 RETURNING user_id, access_id
`

type DeleteUserAccessParams struct {
	UserID   uuid.UUID `db:"user_id" json:"user_id"`
	AccessID uuid.UUID `db:"access_id" json:"access_id"`
}

func (q *Queries) DeleteUserAccess(ctx context.Context, arg *DeleteUserAccessParams) (*UserAccess, error) {
	row := q.db.QueryRow(ctx, DeleteUserAccess, arg.UserID, arg.AccessID)
	var i UserAccess
	err := row.Scan(&i.UserID, &i.AccessID)
	return &i, err
}

const FindUserAccess = `-- name: FindUserAccess :many
SELECT user_id, access_id FROM "user_access" WHERE "user_id" = $1 OR "access_id" = $2
`

type FindUserAccessParams struct {
	UserID   uuid.UUID `db:"user_id" json:"user_id"`
	AccessID uuid.UUID `db:"access_id" json:"access_id"`
}

func (q *Queries) FindUserAccess(ctx context.Context, arg *FindUserAccessParams) ([]*UserAccess, error) {
	rows, err := q.db.Query(ctx, FindUserAccess, arg.UserID, arg.AccessID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []*UserAccess{}
	for rows.Next() {
		var i UserAccess
		if err := rows.Scan(&i.UserID, &i.AccessID); err != nil {
			return nil, err
		}
		items = append(items, &i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const FindUserAccessPage = `-- name: FindUserAccessPage :many
SELECT user_id, access_id FROM "user_access" WHERE "user_id" = $1 OR "access_id" = $2 LIMIT $3 OFFSET $4
`

type FindUserAccessPageParams struct {
	UserID   uuid.UUID `db:"user_id" json:"user_id"`
	AccessID uuid.UUID `db:"access_id" json:"access_id"`
	Limit    int32     `db:"limit" json:"limit"`
	Offset   int32     `db:"offset" json:"offset"`
}

func (q *Queries) FindUserAccessPage(ctx context.Context, arg *FindUserAccessPageParams) ([]*UserAccess, error) {
	rows, err := q.db.Query(ctx, FindUserAccessPage,
		arg.UserID,
		arg.AccessID,
		arg.Limit,
		arg.Offset,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []*UserAccess{}
	for rows.Next() {
		var i UserAccess
		if err := rows.Scan(&i.UserID, &i.AccessID); err != nil {
			return nil, err
		}
		items = append(items, &i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const GetUserAccess = `-- name: GetUserAccess :one
SELECT user_id, access_id FROM "user_access" WHERE "user_id" = $1 AND "access_id" = $2 LIMIT 1
`

type GetUserAccessParams struct {
	UserID   uuid.UUID `db:"user_id" json:"user_id"`
	AccessID uuid.UUID `db:"access_id" json:"access_id"`
}

func (q *Queries) GetUserAccess(ctx context.Context, arg *GetUserAccessParams) (*UserAccess, error) {
	row := q.db.QueryRow(ctx, GetUserAccess, arg.UserID, arg.AccessID)
	var i UserAccess
	err := row.Scan(&i.UserID, &i.AccessID)
	return &i, err
}

const NewUserAccess = `-- name: NewUserAccess :one
INSERT INTO "user_access" ("user_id", "access_id") VALUES ($1, $2) RETURNING user_id, access_id
`

type NewUserAccessParams struct {
	UserID   uuid.UUID `db:"user_id" json:"user_id"`
	AccessID uuid.UUID `db:"access_id" json:"access_id"`
}

func (q *Queries) NewUserAccess(ctx context.Context, arg *NewUserAccessParams) (*UserAccess, error) {
	row := q.db.QueryRow(ctx, NewUserAccess, arg.UserID, arg.AccessID)
	var i UserAccess
	err := row.Scan(&i.UserID, &i.AccessID)
	return &i, err
}