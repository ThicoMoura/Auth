package repository

import (
	"context"

	"github.com/ThicoMoura/Auth/db/model"
	db "github.com/ThicoMoura/Auth/db/sqlc"
)

type IRepository interface {
	New(ctx context.Context, req model.Value) (model.Value, error)
	Get(ctx context.Context, req model.Value) (model.Value, error)
	Find(ctx context.Context, req model.Value) ([]model.Value, error)
	List(ctx context.Context, req model.Value) ([]model.Value, error)
	Update(ctx context.Context, req model.Value) (model.Value, error)
	Delete(ctx context.Context, req model.Value) (model.Value, error)
}

type Repository struct {
	store *db.Store
}

func NewRepository(store *db.Store) *Repository {
	return &Repository{
		store: store,
	}
}

func (repo Repository) Table(name string) IRepository {
	switch name {
	case "access":
		return newAccess(repo.store)
	case "group":
		return newGroup(repo.store)
	case "session":
		return newSession(repo.store)
	case "system":
		return newSystem(repo.store)
	case "user":
		return newUser(repo.store)
	default:
		return nil
	}
}
