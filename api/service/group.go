package service

import (
	"context"

	"github.com/ThicoMoura/Auth/api/model"
	"github.com/ThicoMoura/Auth/db/repository"
)

type Group struct {
	repository repository.IRepository
}

func (service Group) New(ctx context.Context, req model.Model) (model.Model, error) {
	return nil, nil
}

func (service Group) Get(ctx context.Context, req model.Model) (model.Model, error) {
	return nil, nil
}

func (service Group) Find(ctx context.Context, req model.Model) ([]model.Model, error) {
	return nil, nil
}

func (service Group) List(ctx context.Context, req model.Model) ([]model.Model, error) {
	return nil, nil
}

func (service Group) Update(ctx context.Context, req model.Model) (model.Model, error) {
	return nil, nil
}

func (service Group) Delete(ctx context.Context, req model.Model) (model.Model, error) {
	return nil, nil
}

func NewGroup(repository repository.IRepository) Service {
	return &Group{
		repository: repository,
	}
}
