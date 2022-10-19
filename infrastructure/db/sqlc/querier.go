// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.15.0

package db

import (
	"context"
)

type Querier interface {
	CreateUsers(ctx context.Context, arg CreateUsersParams) (User, error)
	DeleteUsers(ctx context.Context, id int64) error
	GetUsers(ctx context.Context, id int64) (User, error)
	ListUsers(ctx context.Context) ([]User, error)
	UpdateUsers(ctx context.Context, arg UpdateUsersParams) error
}

var _ Querier = (*Queries)(nil)
