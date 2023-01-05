package service

import (
	"context"

	"github.com/ThicoMoura/Auth/api/model"
)

type Service interface {
	New(ctx context.Context, req model.Model) (model.Model, error)
	Get(ctx context.Context, req model.Model) (model.Model, error)
	Find(ctx context.Context, req model.Model) ([]model.Model, error)
	List(ctx context.Context, req model.Model) ([]model.Model, error)
	Update(ctx context.Context, req model.Model) (model.Model, error)
	Delete(ctx context.Context, req model.Model) (model.Model, error)
}
