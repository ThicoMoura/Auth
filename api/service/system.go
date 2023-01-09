package service

import (
	"context"

	"github.com/ThicoMoura/Auth/api/model"
	"github.com/ThicoMoura/Auth/db/repository"
)

type System struct {
	repository repository.IRepository
}

func (service System) New(ctx context.Context, req model.Model) (model.Model, error) {
	return nil, nil
}

func (service System) Get(ctx context.Context, req model.Model) (model.Model, error) {
	return nil, nil
}

func (service System) Find(ctx context.Context, req model.Model) ([]model.Model, error) {
	return nil, nil
}

func (service System) List(ctx context.Context, req model.Model) ([]model.Model, error) {
	return nil, nil
}

func (service System) Update(ctx context.Context, req model.Model) (model.Model, error) {
	return nil, nil
}

func (service System) Delete(ctx context.Context, req model.Model) (model.Model, error) {
	return nil, nil
}

func NewSystem(repository repository.IRepository) Service {
	return &System{
		repository: repository,
	}
}
