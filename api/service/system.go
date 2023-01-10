package service

import (
	"context"

	"github.com/ThicoMoura/Auth/api/model"
	"github.com/ThicoMoura/Auth/db/repository"
	"github.com/google/uuid"
)

type System struct {
	repository repository.IRepository
}

func (service System) New(ctx context.Context, req model.Model) (model.Model, error) {
	system, err := service.repository.New(ctx, req)
	if err != nil {
		return nil, err
	}

	return model.NewSystem(system.Get("ID").(uuid.UUID), system.Get("Name").(string), nil, system.Get("Active").(bool)), nil
}

func (service System) Get(ctx context.Context, req model.Model) (model.Model, error) {
	system, err := service.repository.Get(ctx, req)
	if err != nil {
		return nil, err
	}

	return model.NewSystem(system.Get("ID").(uuid.UUID), system.Get("Name").(string), nil, system.Get("Active").(bool)), nil
}

func (service System) Find(ctx context.Context, req model.Model) ([]model.Model, error) {
	res, err := service.repository.Find(ctx, req)
	if err != nil {
		return nil, err
	}

	var systems []model.Model

	for _, system := range res {
		systems = append(systems, model.NewSystem(system.Get("ID").(uuid.UUID), system.Get("Name").(string), nil, system.Get("Active").(bool)))
	}

	return systems, nil
}

func (service System) List(ctx context.Context, req model.Model) ([]model.Model, error) {
	res, err := service.repository.List(ctx, req)
	if err != nil {
		return nil, err
	}

	var list []model.Model

	for _, system := range res {
		list = append(list, model.NewSystem(system.Get("ID").(uuid.UUID), system.Get("Name").(string), nil, system.Get("Active").(bool)))
	}

	return list, nil
}

func (service System) Update(ctx context.Context, req model.Model) (model.Model, error) {
	system, err := service.repository.Update(ctx, req)
	if err != nil {
		return nil, err
	}

	return model.NewSystem(system.Get("ID").(uuid.UUID), system.Get("Name").(string), nil, system.Get("Active").(bool)), nil
}

func (service System) Delete(ctx context.Context, req model.Model) (model.Model, error) {
	system, err := service.repository.Delete(ctx, req)
	if err != nil {
		return nil, err
	}

	return model.NewSystem(system.Get("ID").(uuid.UUID), system.Get("Name").(string), nil, system.Get("Active").(bool)), nil
}

func NewSystem(repository repository.IRepository) Service {
	return &System{
		repository: repository,
	}
}
