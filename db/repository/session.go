package repository

import (
	"context"

	"github.com/ThicoMoura/Auth/db/model"
	db "github.com/ThicoMoura/Auth/db/sqlc"
)

type session struct {
	store *db.Store
}

func newSession(store *db.Store) *session {
	return &session{
		store: store,
	}
}

func (repository session) New(ctx context.Context, req model.Value) (model.Value, error) {
	return nil, nil
}

func (repository session) Get(ctx context.Context, req model.Value) (model.Value, error) {
	return nil, nil
}

func (repository session) Find(ctx context.Context, req model.Value) ([]model.Value, error) {
	return nil, nil
}

func (repository session) List(ctx context.Context, req model.Value) ([]model.Value, error) {
	return nil, nil
}

func (repository session) Update(ctx context.Context, req model.Value) (model.Value, error) {
	return nil, nil
}

func (repository session) Delete(ctx context.Context, req model.Value) (model.Value, error) {
	return nil, nil
}
