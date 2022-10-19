// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.15.0
// source: users.sql

package db

import (
	"context"
)

const createUsers = `-- name: CreateUsers :one
INSERT INTO users (cpf, name, pass) VALUES ($1, $2, $3) RETURNING id, cpf, name, pass, active
`

type CreateUsersParams struct {
	Cpf  string `json:"cpf"`
	Name string `json:"name"`
	Pass string `json:"pass"`
}

func (q *Queries) CreateUsers(ctx context.Context, arg CreateUsersParams) (User, error) {
	row := q.queryRow(ctx, q.createUsersStmt, createUsers, arg.Cpf, arg.Name, arg.Pass)
	var i User
	err := row.Scan(
		&i.ID,
		&i.Cpf,
		&i.Name,
		&i.Pass,
		&i.Active,
	)
	return i, err
}

const deleteUsers = `-- name: DeleteUsers :exec
DELETE FROM users WHERE id = $1
`

func (q *Queries) DeleteUsers(ctx context.Context, id int64) error {
	_, err := q.exec(ctx, q.deleteUsersStmt, deleteUsers, id)
	return err
}

const getUsers = `-- name: GetUsers :one
SELECT id, cpf, name, pass, active FROM users WHERE id = $1 LIMIT 1
`

func (q *Queries) GetUsers(ctx context.Context, id int64) (User, error) {
	row := q.queryRow(ctx, q.getUsersStmt, getUsers, id)
	var i User
	err := row.Scan(
		&i.ID,
		&i.Cpf,
		&i.Name,
		&i.Pass,
		&i.Active,
	)
	return i, err
}

const listUsers = `-- name: ListUsers :many
SELECT id, cpf, name, pass, active FROM users
`

func (q *Queries) ListUsers(ctx context.Context) ([]User, error) {
	rows, err := q.query(ctx, q.listUsersStmt, listUsers)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []User
	for rows.Next() {
		var i User
		if err := rows.Scan(
			&i.ID,
			&i.Cpf,
			&i.Name,
			&i.Pass,
			&i.Active,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const updateUsers = `-- name: UpdateUsers :exec
UPDATE users SET cpf = $2, name = $3, pass = $4 WHERE id = $1
`

type UpdateUsersParams struct {
	ID   int64  `json:"id"`
	Cpf  string `json:"cpf"`
	Name string `json:"name"`
	Pass string `json:"pass"`
}

func (q *Queries) UpdateUsers(ctx context.Context, arg UpdateUsersParams) error {
	_, err := q.exec(ctx, q.updateUsersStmt, updateUsers,
		arg.ID,
		arg.Cpf,
		arg.Name,
		arg.Pass,
	)
	return err
}