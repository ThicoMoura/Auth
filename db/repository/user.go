package repository

import (
	"context"

	"github.com/ThicoMoura/Auth/db/model"
	db "github.com/ThicoMoura/Auth/db/sqlc"
	"github.com/google/uuid"
)

type user struct {
	store *db.Store
}

func newUser(store *db.Store) *user {
	return &user{
		store: store,
	}
}

func (repository user) New(ctx context.Context, req model.Value) (model.Value, error) {
	user, err := repository.store.NewUser(ctx, &db.NewUserParams{
		Group: req.Get("Group").(uuid.UUID),
		Email: req.Get("Email").(string),
		Name:  req.Get("Name").(string),
		Pass:  req.Get("Pass").(string),
	})
	if err != nil {
		return nil, err
	}

	return model.New(map[string]interface{}{
		"ID":     user.ID,
		"Group":  user.Group,
		"Email":  user.Email,
		"Name":   user.Name,
		"Pass":   user.Pass,
		"Active": user.Active,
	}), nil
}

func (repository user) Get(ctx context.Context, req model.Value) (model.Value, error) {
	arg := &db.GetUserParams{}

	if id, ok := req.Get("ID").(uuid.UUID); ok {
		arg.ID = id
	} else if email, ok := req.Get("Email").(string); ok {
		arg.Email = email
	}

	user, err := repository.store.GetUser(ctx, arg)
	if err != nil {
		return nil, err
	}

	return model.New(map[string]interface{}{
		"ID":     user.ID,
		"Group":  user.Group,
		"Email":  user.Email,
		"Name":   user.Name,
		"Pass":   user.Pass,
		"Active": user.Active,
	}), nil
}

func (repository user) Find(ctx context.Context, req model.Value) ([]model.Value, error) {
	_, limit := req.Get("Limit").(int32)
	_, offset := req.Get("Offset").(int32)

	if limit && offset {
		arg := &db.FindUserPageParams{
			Limit:  req.Get("Limit").(int32),
			Offset: req.Get("Offset").(int32),
		}

		if group, ok := req.Get("Group").(uuid.UUID); ok {
			arg.Group = group
		}

		if name, ok := req.Get("Name").(string); ok {
			arg.Name = name
		}

		res, err := repository.store.FindUserPage(ctx, arg)
		if err != nil {
			return nil, err
		}

		var users []model.Value

		for _, user := range res {
			users = append(users, model.New(map[string]interface{}{
				"ID":     user.ID,
				"Group":  user.Group,
				"Email":  user.Email,
				"Name":   user.Name,
				"Pass":   user.Pass,
				"Active": user.Active,
			}))
		}

		return users, nil
	}

	arg := &db.FindUserParams{}

	if group, ok := req.Get("Group").(uuid.UUID); ok {
		arg.Group = group
	}

	if name, ok := req.Get("Name").(string); ok {
		arg.Name = name
	}

	res, err := repository.store.FindUser(ctx, arg)
	if err != nil {
		return nil, err
	}

	var users []model.Value

	for _, user := range res {
		users = append(users, model.New(map[string]interface{}{
			"ID":     user.ID,
			"Group":  user.Group,
			"Email":  user.Email,
			"Name":   user.Name,
			"Pass":   user.Pass,
			"Active": user.Active,
		}))
	}

	return users, nil
}

func (repository user) List(ctx context.Context, req model.Value) ([]model.Value, error) {
	_, limit := req.Get("Limit").(int32)
	_, offset := req.Get("Offset").(int32)

	if limit && offset {
		arg := &db.ListUserPageParams{
			Limit:  req.Get("Limit").(int32),
			Offset: req.Get("Offset").(int32),
		}
		res, err := repository.store.ListUserPage(ctx, arg)
		if err != nil {
			return nil, err
		}

		var users []model.Value

		for _, user := range res {
			users = append(users, model.New(map[string]interface{}{
				"ID":     user.ID,
				"Group":  user.Group,
				"Email":  user.Email,
				"Name":   user.Name,
				"Pass":   user.Pass,
				"Active": user.Active,
			}))
		}

		return users, nil
	}

	res, err := repository.store.ListUser(ctx)
	if err != nil {
		return nil, err
	}

	var users []model.Value

	for _, user := range res {
		users = append(users, model.New(map[string]interface{}{
			"ID":     user.ID,
			"Group":  user.Group,
			"Email":  user.Email,
			"Name":   user.Name,
			"Pass":   user.Pass,
			"Active": user.Active,
		}))
	}

	return users, nil
}

func (repository user) Update(ctx context.Context, req model.Value) (res model.Value, err error) {
	err = repository.store.TX(ctx, func(q *db.Queries) error {
		var user *db.User
		var err error

		arg := &db.UpdateUserParams{
			ID: req.Get("ID").(uuid.UUID),
		}

		if name, ok := req.Get("Name").(string); ok {
			arg.Name = name
		}

		if pass, ok := req.Get("Pass").(string); ok {
			arg.Pass = pass
		}

		user, err = repository.store.UpdateUser(ctx, arg)
		if err != nil {
			return err
		}

		if active, ok := req.Get("Active").(bool); ok {
			user, err = repository.store.UpdateUserActive(ctx, &db.UpdateUserActiveParams{
				ID:     req.Get("ID").(uuid.UUID),
				Active: active,
			})
			if err != nil {
				return err
			}
		}

		res = model.New(map[string]interface{}{
			"ID":     user.ID,
			"Group":  user.Group,
			"Email":  user.Email,
			"Name":   user.Name,
			"Pass":   user.Pass,
			"Active": user.Active,
		})

		return nil
	})
	return
}

func (repository user) Delete(ctx context.Context, req model.Value) (model.Value, error) {
	user, err := repository.store.DeleteUser(ctx, req.Get("ID").(uuid.UUID))
	if err != nil {
		return nil, err
	}

	return model.New(map[string]interface{}{
		"ID":     user.ID,
		"Group":  user.Group,
		"Email":  user.Email,
		"Name":   user.Name,
		"Pass":   user.Pass,
		"Active": user.Active,
	}), nil
}
