package repository

import (
	"context"

	"github.com/ThicoMoura/Auth/db/model"
	db "github.com/ThicoMoura/Auth/db/sqlc"
)

type user struct {
	store *db.Store
}

func newUser(store *db.Store) *user {
	return &user{
		store: store,
	}
}

func (repository user) New(ctx context.Context, req model.Value) (model.Value, error) {
	return nil, nil
}

func (repository user) Get(ctx context.Context, req model.Value) (model.Value, error) {
	return nil, nil
}

func (repository user) Find(ctx context.Context, req model.Value) ([]model.Value, error) {
	return nil, nil
}

func (repository user) List(ctx context.Context, req model.Value) ([]model.Value, error) {
	return nil, nil
}

func (repository user) Update(ctx context.Context, req model.Value) (model.Value, error) {
	return nil, nil
}

func (repository user) Delete(ctx context.Context, req model.Value) (model.Value, error) {
	return nil, nil
}
