package repository

import (
	"context"

	"github.com/ThicoMoura/Auth/db/model"
	db "github.com/ThicoMoura/Auth/db/sqlc"
	"github.com/google/uuid"
)

type group struct {
	store *db.Store
}

func newGroup(store *db.Store) *group {
	return &group{
		store: store,
	}
}

func (repository group) New(ctx context.Context, req model.Value) (model.Value, error) {
	group, err := repository.store.NewGroup(ctx, req.Get("Name").(string))
	if err != nil {
		return nil, err
	}

	return model.New(map[string]interface{}{
		"ID":     group.ID,
		"Name":   group.Name,
		"Active": group.Active,
	}), nil
}

func (repository group) Get(ctx context.Context, req model.Value) (model.Value, error) {
	group, err := repository.store.GetGroup(ctx, req.Get("ID").(uuid.UUID))
	if err != nil {
		return nil, err
	}

	return model.New(map[string]interface{}{
		"ID":     group.ID,
		"Name":   group.Name,
		"Active": group.Active,
	}), nil
}

func (repository group) Find(ctx context.Context, req model.Value) ([]model.Value, error) {
	_, okLimit := req.Get("Limit").(int32)
	_, okOffset := req.Get("Offset").(int32)

	if okLimit && okOffset {
		res, err := repository.store.FindGroupPage(ctx, &db.FindGroupPageParams{
			Name:   req.Get("Name").(string),
			Limit:  req.Get("Limit").(int32),
			Offset: req.Get("Offset").(int32),
		})
		if err != nil {
			return nil, err
		}

		var groups []model.Value

		for _, group := range res {
			groups = append(groups, model.New(map[string]interface{}{
				"ID":     group.ID,
				"Name":   group.Name,
				"Active": group.Active,
			}))
		}

		return groups, nil
	}

	res, err := repository.store.FindGroup(ctx, req.Get("Name").(string))
	if err != nil {
		return nil, err
	}

	var groups []model.Value

	for _, group := range res {
		groups = append(groups, model.New(map[string]interface{}{
			"ID":     group.ID,
			"Name":   group.Name,
			"Active": group.Active,
		}))
	}

	return groups, nil
}

func (repository group) List(ctx context.Context, req model.Value) ([]model.Value, error) {
	_, okLimit := req.Get("Limit").(int32)
	_, okOffset := req.Get("Offset").(int32)

	if okLimit && okOffset {
		res, err := repository.store.ListGroupPage(ctx, &db.ListGroupPageParams{
			Limit:  req.Get("Limit").(int32),
			Offset: req.Get("Offset").(int32),
		})
		if err != nil {
			return nil, err
		}

		var groups []model.Value

		for _, group := range res {
			groups = append(groups, model.New(map[string]interface{}{
				"ID":     group.ID,
				"Name":   group.Name,
				"Active": group.Active,
			}))
		}

		return groups, nil
	}

	res, err := repository.store.ListGroup(ctx)
	if err != nil {
		return nil, err
	}

	var groups []model.Value

	for _, group := range res {
		groups = append(groups, model.New(map[string]interface{}{
			"ID":     group.ID,
			"Name":   group.Name,
			"Active": group.Active,
		}))
	}

	return groups, nil
}

func (repository group) Update(ctx context.Context, req model.Value) (res model.Value, err error) {
	err = repository.store.TX(ctx, func(q *db.Queries) error {
		var group *db.Group
		var err error

		if name, ok := req.Get("Name").(string); ok {
			group, err = repository.store.UpdateGroup(ctx, &db.UpdateGroupParams{
				ID:   req.Get("ID").(uuid.UUID),
				Name: name,
			})
			if err != nil {
				return err
			}
		}

		if active, ok := req.Get("Active").(bool); ok {
			group, err = repository.store.UpdateGroupActive(ctx, &db.UpdateGroupActiveParams{
				ID:     req.Get("ID").(uuid.UUID),
				Active: active,
			})
			if err != nil {
				return err
			}
		}

		res = model.New(map[string]interface{}{
			"ID":     group.ID,
			"Name":   group.Name,
			"Active": group.Active,
		})

		return nil
	})
	if err != nil {
		return
	}

	return
}

func (repository group) Delete(ctx context.Context, req model.Value) (model.Value, error) {
	group, err := repository.store.DeleteGroup(ctx, req.Get("ID").(uuid.UUID))
	if err != nil {
		return nil, err
	}

	return model.New(map[string]interface{}{
		"ID":     group.ID,
		"Name":   group.Name,
		"Active": group.Active,
	}), nil
}
