package service

import (
	"context"

	"github.com/ThicoMoura/Auth/api/model"
	"github.com/ThicoMoura/Auth/db/repository"
)

type User struct {
	repository repository.IRepository
}

func (service User) New(ctx context.Context, req model.Model) (model.Model, error) {
	return nil, nil
}

func (service User) Get(ctx context.Context, req model.Model) (model.Model, error) {
	return nil, nil
}

func (service User) Find(ctx context.Context, req model.Model) ([]model.Model, error) {
	return nil, nil
}

func (service User) List(ctx context.Context, req model.Model) ([]model.Model, error) {
	return nil, nil
}

func (service User) Update(ctx context.Context, req model.Model) (model.Model, error) {
	return nil, nil
}

func (service User) Delete(ctx context.Context, req model.Model) (model.Model, error) {
	return nil, nil
}

func NewUser(repository repository.IRepository) Service {
	return &User{
		repository: repository,
	}
}
