package repository

import (
	"context"

	"github.com/ThicoMoura/Auth/db/model"
	db "github.com/ThicoMoura/Auth/db/sqlc"
)

type group struct {
	store *db.Store
}

func newGroup(store *db.Store) *group {
	return &group{
		store: store,
	}
}

func (repository group) New(ctx context.Context, req model.Value) (model.Value, error) {
	return nil, nil
}

func (repository group) Get(ctx context.Context, req model.Value) (model.Value, error) {
	return nil, nil
}

func (repository group) Find(ctx context.Context, req model.Value) ([]model.Value, error) {
	return nil, nil
}

func (repository group) List(ctx context.Context, req model.Value) ([]model.Value, error) {
	return nil, nil
}

func (repository group) Update(ctx context.Context, req model.Value) (model.Value, error) {
	return nil, nil
}

func (repository group) Delete(ctx context.Context, req model.Value) (model.Value, error) {
	return nil, nil
}
