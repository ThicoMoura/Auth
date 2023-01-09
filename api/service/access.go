package service

import (
	"context"

	"github.com/ThicoMoura/Auth/api/model"
	"github.com/ThicoMoura/Auth/db/repository"
	"github.com/google/uuid"
)

type Access struct {
	repository repository.IRepository
}

func (service Access) New(ctx context.Context, req model.Model) (model.Model, error) {
	return nil, nil
}

func (service Access) Get(ctx context.Context, req model.Model) (model.Model, error) {
	res, err := service.repository.Get(ctx, req)
	if err != nil {
		return nil, err
	}

	return model.NewAccess(res.Get("ID").(uuid.UUID), res.Get("System").(uuid.UUID), res.Get("Table").(string), res.Get("Type").([]string)), nil
}

func (service Access) Find(ctx context.Context, req model.Model) ([]model.Model, error) {
	return nil, nil
}

func (service Access) List(ctx context.Context, req model.Model) ([]model.Model, error) {
	return nil, nil
}

func (service Access) Update(ctx context.Context, req model.Model) (model.Model, error) {
	return nil, nil
}

func (service Access) Delete(ctx context.Context, req model.Model) (model.Model, error) {
	return nil, nil
}

func NewAccess(repository repository.IRepository) Service {
	return &Access{
		repository: repository,
	}
}
