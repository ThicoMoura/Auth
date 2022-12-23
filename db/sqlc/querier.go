// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.16.0

package db

import (
	"context"

	"github.com/google/uuid"
)

type Querier interface {
	DeleteGroup(ctx context.Context, db DBTX, id uuid.UUID) (*Group, error)
	DeleteUser(ctx context.Context, db DBTX, id uuid.UUID) (*User, error)
	GetGroup(ctx context.Context, db DBTX, id uuid.UUID) (*Group, error)
	GetUser(ctx context.Context, db DBTX, id uuid.UUID) (*User, error)
	GetUserByCPF(ctx context.Context, db DBTX, cpf string) (*User, error)
	GetUserByGroup(ctx context.Context, db DBTX, group uuid.UUID) ([]*User, error)
	GetUserByGroupPage(ctx context.Context, db DBTX, arg *GetUserByGroupPageParams) ([]*User, error)
	ListGroup(ctx context.Context, db DBTX) ([]*Group, error)
	ListGroupPage(ctx context.Context, db DBTX, arg *ListGroupPageParams) ([]*Group, error)
	ListUser(ctx context.Context, db DBTX) ([]*User, error)
	ListUserPage(ctx context.Context, db DBTX, arg *ListUserPageParams) ([]*User, error)
	NewGroup(ctx context.Context, db DBTX, name string) (*Group, error)
	NewUser(ctx context.Context, db DBTX, arg *NewUserParams) (*User, error)
	UpdateGroup(ctx context.Context, db DBTX, arg *UpdateGroupParams) (*Group, error)
	UpdateUser(ctx context.Context, db DBTX, arg *UpdateUserParams) (*User, error)
	UpdateUserPass(ctx context.Context, db DBTX, arg *UpdateUserPassParams) (*User, error)
}

var _ Querier = (*Queries)(nil)
