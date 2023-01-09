package service

import (
	"context"

	"github.com/ThicoMoura/Auth/api/model"
	"github.com/ThicoMoura/Auth/db/repository"
)

type Session struct {
	repository repository.IRepository
}

func (service Session) New(ctx context.Context, req model.Model) (model.Model, error) {
	return nil, nil
}

func (service Session) Get(ctx context.Context, req model.Model) (model.Model, error) {
	return nil, nil
}

func (service Session) Find(ctx context.Context, req model.Model) ([]model.Model, error) {
	return nil, nil
}

func (service Session) List(ctx context.Context, req model.Model) ([]model.Model, error) {
	return nil, nil
}

func (service Session) Update(ctx context.Context, req model.Model) (model.Model, error) {
	return nil, nil
}

func (service Session) Delete(ctx context.Context, req model.Model) (model.Model, error) {
	return nil, nil
}

func NewSession(repository repository.IRepository) Service {
	return &Session{
		repository: repository,
	}
}
