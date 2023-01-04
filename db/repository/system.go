package repository

import (
	"context"

	"github.com/ThicoMoura/Auth/db/model"
	db "github.com/ThicoMoura/Auth/db/sqlc"
	"github.com/google/uuid"
)

type system struct {
	store *db.Store
}

func newSystem(store *db.Store) *system {
	return &system{
		store: store,
	}
}

func (repository system) New(ctx context.Context, req model.Value) (model.Value, error) {
	system, err := repository.store.NewSystem(ctx, &db.NewSystemParams{
		Name:   req.Get("Name").(string),
		Tables: req.Get("Tables").([]string),
	})
	if err != nil {
		return nil, err
	}

	return model.New(map[string]interface{}{
		"ID":     system.ID,
		"Name":   system.Name,
		"Tables": system.Tables,
		"Active": system.Active,
	}), nil
}

func (repository system) Get(ctx context.Context, req model.Value) (model.Value, error) {
	system, err := repository.store.GetSystem(ctx, req.Get("ID").(uuid.UUID))
	if err != nil {
		return nil, err
	}

	return model.New(map[string]interface{}{
		"ID":     system.ID,
		"Name":   system.Name,
		"Tables": system.Tables,
		"Active": system.Active,
	}), nil
}

func (repository system) Find(ctx context.Context, req model.Value) ([]model.Value, error) {
	_, okOffset := req.Get("Offset").(int32)
	_, okLimit := req.Get("Limit").(int32)
	if okOffset && okLimit {
		systems, err := repository.store.FindSystemPage(ctx, &db.FindSystemPageParams{
			Name:   req.Get("Name").(string),
			Limit:  req.Get("Limit").(int32),
			Offset: req.Get("Offset").(int32),
		})
		if err != nil {
			return nil, err
		}

		var res []model.Value

		for _, system := range systems {
			res = append(res, model.New(map[string]interface{}{
				"ID":     system.ID,
				"Name":   system.Name,
				"Tables": system.Tables,
				"Active": system.Active,
			}))
		}

		return res, nil
	}

	systems, err := repository.store.FindSystem(ctx, req.Get("Name").(string))
	if err != nil {
		return nil, err
	}

	var res []model.Value

	for _, system := range systems {
		res = append(res, model.New(map[string]interface{}{
			"ID":     system.ID,
			"Name":   system.Name,
			"Tables": system.Tables,
			"Active": system.Active,
		}))
	}

	return res, nil
}

func (repository system) List(ctx context.Context, req model.Value) ([]model.Value, error) {
	_, okOffset := req.Get("Offset").(int32)
	_, okLimit := req.Get("Limit").(int32)
	if okOffset && okLimit {
		systems, err := repository.store.ListSystemPage(ctx, &db.ListSystemPageParams{
			Limit:  req.Get("Limit").(int32),
			Offset: req.Get("Offset").(int32),
		})
		if err != nil {
			return nil, err
		}

		var res []model.Value

		for _, system := range systems {
			res = append(res, model.New(map[string]interface{}{
				"ID":     system.ID,
				"Name":   system.Name,
				"Tables": system.Tables,
				"Active": system.Active,
			}))
		}

		return res, nil
	}

	systems, err := repository.store.ListSystem(ctx)
	if err != nil {
		return nil, err
	}

	var res []model.Value

	for _, system := range systems {
		res = append(res, model.New(map[string]interface{}{
			"ID":     system.ID,
			"Name":   system.Name,
			"Tables": system.Tables,
			"Active": system.Active,
		}))
	}

	return res, nil
}

func (repository system) Update(ctx context.Context, req model.Value) (res model.Value, err error) {
	err = repository.store.TX(ctx, func(q *db.Queries) error {
		arg := &db.UpdateSystemParams{
			ID: req.Get("ID").(uuid.UUID),
		}

		if name, ok := req.Get("Name").(string); ok {
			arg.Name = name
		}

		if tables, ok := req.Get("Tables").([]string); ok {
			arg.Tables = tables
		}

		system, err := q.UpdateSystem(ctx, arg)
		if err != nil {
			return err
		}

		if active, ok := req.Get("Active").(bool); ok {
			system, err = q.UpdateActiveSystem(ctx, &db.UpdateActiveSystemParams{
				ID:     req.Get("ID").(uuid.UUID),
				Active: active,
			})

			if err != nil {
				return err
			}
		}

		res = model.New(map[string]interface{}{
			"ID":     system.ID,
			"Name":   system.Name,
			"Tables": system.Tables,
			"Active": system.Active,
		})

		return nil
	})
	return
}

func (repository system) Delete(ctx context.Context, req model.Value) (model.Value, error) {
	system, err := repository.store.DeleteSystem(ctx, req.Get("ID").(uuid.UUID))
	if err != nil {
		return nil, err
	}

	return model.New(map[string]interface{}{
		"ID":     system.ID,
		"Name":   system.Name,
		"Tables": system.Tables,
		"Active": system.Active,
	}), nil
}
