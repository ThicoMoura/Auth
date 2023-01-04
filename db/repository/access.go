package repository

import (
	"context"

	"github.com/ThicoMoura/Auth/db/model"
	db "github.com/ThicoMoura/Auth/db/sqlc"
)

type access struct {
	store *db.Store
}

func newAccess(store *db.Store) *access {
	return &access{
		store: store,
	}
}

func (repository access) New(ctx context.Context, req model.Value) (model.Value, error) {
	return nil, nil
}

func (repository access) Get(ctx context.Context, req model.Value) (model.Value, error) {
	return nil, nil
}

func (repository access) Find(ctx context.Context, req model.Value) ([]model.Value, error) {
	return nil, nil
}

func (repository access) List(ctx context.Context, req model.Value) ([]model.Value, error) {
	return nil, nil
}

func (repository access) Update(ctx context.Context, req model.Value) (model.Value, error) {
	return nil, nil
}

func (repository access) Delete(ctx context.Context, req model.Value) (model.Value, error) {
	return nil, nil
}
