package repository

import (
	"context"

	"github.com/ThicoMoura/Auth/db/model"
	db "github.com/ThicoMoura/Auth/db/sqlc"
	"github.com/google/uuid"
)

type access struct {
	store *db.Store
}

func newAccess(store *db.Store) *access {
	return &access{
		store: store,
	}
}

func (repository access) New(ctx context.Context, req model.Value) (model.Value, error) {
	access, err := repository.store.NewAccess(ctx, &db.NewAccessParams{
		System: req.Get("System").(uuid.UUID),
		Table:  req.Get("Table").(string),
		Type:   req.Get("Type").([]string),
	})
	if err != nil {
		return nil, err
	}

	return model.New(map[string]interface{}{
		"ID":     access.ID,
		"System": access.System,
		"Table":  access.Table,
		"Type":   access.Type,
	}), nil
}

func (repository access) Get(ctx context.Context, req model.Value) (model.Value, error) {
	access, err := repository.store.GetAccess(ctx, req.Get("ID").(uuid.UUID))
	if err != nil {
		return nil, err
	}

	return model.New(map[string]interface{}{
		"ID":     access.ID,
		"System": access.System,
		"Table":  access.Table,
		"Type":   access.Type,
	}), nil
}

func (repository access) Find(ctx context.Context, req model.Value) ([]model.Value, error) {
	_, okLimit := req.Get("Limit").(int32)
	_, okOffset := req.Get("Offset").(int32)

	if okLimit && okOffset {
		arg := &db.FindAccessPageParams{
			Limit:  req.Get("Limit").(int32),
			Offset: req.Get("Offset").(int32),
		}

		if system, ok := req.Get("System").(uuid.UUID); ok {
			arg.System = system
		}

		if table, ok := req.Get("Table").(string); ok {
			arg.Table = table
		}

		res, err := repository.store.FindAccessPage(ctx, arg)
		if err != nil {
			return nil, err
		}

		var access []model.Value

		for _, ac := range res {
			access = append(access, model.New(map[string]interface{}{
				"ID":     ac.ID,
				"System": ac.System,
				"Table":  ac.Table,
				"Type":   ac.Type,
			}))
		}

		return access, nil
	}

	arg := &db.FindAccessParams{}

	if system, ok := req.Get("System").(uuid.UUID); ok {
		arg.System = system
	}

	if table, ok := req.Get("Table").(string); ok {
		arg.Table = table
	}

	res, err := repository.store.FindAccess(ctx, arg)
	if err != nil {
		return nil, err
	}

	var access []model.Value

	for _, ac := range res {
		access = append(access, model.New(map[string]interface{}{
			"ID":     ac.ID,
			"System": ac.System,
			"Table":  ac.Table,
			"Type":   ac.Type,
		}))
	}

	return access, nil
}

func (repository access) List(ctx context.Context, req model.Value) ([]model.Value, error) {
	_, okLimit := req.Get("Limit").(int32)
	_, okOffset := req.Get("Offset").(int32)

	if okLimit && okOffset {
		res, err := repository.store.ListAccessPage(ctx, &db.ListAccessPageParams{
			Limit:  req.Get("Limit").(int32),
			Offset: req.Get("Offset").(int32),
		})
		if err != nil {
			return nil, err
		}

		var access []model.Value

		for _, ac := range res {
			access = append(access, model.New(map[string]interface{}{
				"ID":     ac.ID,
				"System": ac.System,
				"Table":  ac.Table,
				"Type":   ac.Type,
			}))
		}

		return access, nil
	}

	res, err := repository.store.ListAccess(ctx)
	if err != nil {
		return nil, err
	}

	var access []model.Value

	for _, ac := range res {
		access = append(access, model.New(map[string]interface{}{
			"ID":     ac.ID,
			"System": ac.System,
			"Table":  ac.Table,
			"Type":   ac.Type,
		}))
	}

	return access, nil
}

func (repository access) Update(ctx context.Context, req model.Value) (model.Value, error) {
	access, err := repository.store.UpdateAccess(ctx, &db.UpdateAccessParams{
		ID:   req.Get("ID").(uuid.UUID),
		Type: req.Get("Type").([]string),
	})
	if err != nil {
		return nil, err
	}

	return model.New(map[string]interface{}{
		"ID":     access.ID,
		"System": access.System,
		"Table":  access.Table,
		"Type":   access.Type,
	}), nil
}

func (repository access) Delete(ctx context.Context, req model.Value) (model.Value, error) {
	access, err := repository.store.DeleteAccess(ctx, req.Get("ID").(uuid.UUID))
	if err != nil {
		return nil, err
	}

	return model.New(map[string]interface{}{
		"ID":     access.ID,
		"System": access.System,
		"Table":  access.Table,
		"Type":   access.Type,
	}), nil
}
