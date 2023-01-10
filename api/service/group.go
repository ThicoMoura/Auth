package service

import (
	"context"

	"github.com/ThicoMoura/Auth/api/model"
	"github.com/ThicoMoura/Auth/db/repository"
	"github.com/google/uuid"
)

type Group struct {
	repository repository.IRepository
}

func (service Group) New(ctx context.Context, req model.Model) (model.Model, error) {
	group, err := service.repository.New(ctx, req)
	if err != nil {
		return nil, err
	}

	return model.NewGroup(group.Get("ID").(uuid.UUID), group.Get("Name").(string), nil, nil, group.Get("Active").(bool)), nil
}

func (service Group) Get(ctx context.Context, req model.Model) (model.Model, error) {
	group, err := service.repository.Get(ctx, req)
	if err != nil {
		return nil, err
	}

	return model.NewGroup(group.Get("ID").(uuid.UUID), group.Get("Name").(string), nil, nil, group.Get("Active").(bool)), nil
}

func (service Group) Find(ctx context.Context, req model.Model) ([]model.Model, error) {
	res, err := service.repository.Find(ctx, req)
	if err != nil {
		return nil, err
	}

	var groups []model.Model

	for _, group := range res {
		groups = append(groups, model.NewGroup(group.Get("ID").(uuid.UUID), group.Get("Name").(string), nil, nil, group.Get("Active").(bool)))
	}

	return groups, nil
}

func (service Group) List(ctx context.Context, req model.Model) ([]model.Model, error) {
	res, err := service.repository.List(ctx, req)
	if err != nil {
		return nil, err
	}

	var list []model.Model

	for _, group := range res {
		list = append(list, model.NewGroup(group.Get("ID").(uuid.UUID), group.Get("Name").(string), nil, nil, group.Get("Active").(bool)))
	}

	return list, nil
}

func (service Group) Update(ctx context.Context, req model.Model) (model.Model, error) {
	group, err := service.repository.Update(ctx, req)
	if err != nil {
		return nil, err
	}

	return model.NewGroup(group.Get("ID").(uuid.UUID), group.Get("Name").(string), nil, nil, group.Get("Active").(bool)), nil
}

func (service Group) Delete(ctx context.Context, req model.Model) (model.Model, error) {
	group, err := service.repository.Delete(ctx, req)
	if err != nil {
		return nil, err
	}

	return model.NewGroup(group.Get("ID").(uuid.UUID), group.Get("Name").(string), nil, nil, group.Get("Active").(bool)), nil
}

func NewGroup(repository repository.IRepository) Service {
	return &Group{
		repository: repository,
	}
}
