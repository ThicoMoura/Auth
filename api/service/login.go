package service

import (
	"context"

	"github.com/ThicoMoura/Auth/api/model"
	md "github.com/ThicoMoura/Auth/db/model"
	"github.com/ThicoMoura/Auth/db/repository"
	"github.com/google/uuid"
)

type login struct {
	store map[string]repository.IRepository
}

func (service login) New(ctx context.Context, req model.Model) (model.Model, error) {
	return nil, nil
}

func (service login) Get(ctx context.Context, req model.Model) (model.Model, error) {
	user, err := service.store["user"].Get(ctx, md.New(map[string]interface{}{
		"Email": req.Get("Email"),
	}))
	if err != nil {
		return nil, err
	}

	return model.User{
		GroupID: user.Get("Group").(uuid.UUID),
		Email:   user.Get("Email").(string),
		Name:    user.Get("Name").(string),
		Pass:    user.Get("Pass").(string),
		Active:  user.Get("Active").(bool),
	}, nil
}

func (service login) Find(ctx context.Context, req model.Model) ([]model.Model, error) {
	return nil, nil
}

func (service login) List(ctx context.Context, req model.Model) ([]model.Model, error) {
	return nil, nil
}

func (service login) Update(ctx context.Context, req model.Model) (model.Model, error) {
	return nil, nil
}

func (service login) Delete(ctx context.Context, req model.Model) (model.Model, error) {
	return nil, nil
}

func NewLogin(store map[string]repository.IRepository) Service {
	return &login{
		store: store,
	}
}
