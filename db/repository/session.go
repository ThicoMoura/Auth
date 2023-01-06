package repository

import (
	"context"
	"errors"
	"time"

	"github.com/ThicoMoura/Auth/db/model"
	db "github.com/ThicoMoura/Auth/db/sqlc"
	"github.com/google/uuid"
)

type session struct {
	store *db.Store
}

func newSession(store *db.Store) *session {
	return &session{
		store: store,
	}
}

func (repository session) New(ctx context.Context, req model.Value) (model.Value, error) {
	session, err := repository.store.NewSession(ctx, &db.NewSessionParams{
		ID:        req.Get("ID").(uuid.UUID),
		User:      req.Get("User").(uuid.UUID),
		Token:     req.Get("Token").(string),
		Ip:        req.Get("Ip").(string),
		Agent:     req.Get("Agent").(string),
		ExpiresAt: req.Get("ExpiresAt").(time.Time),
	})
	if err != nil {
		return nil, err
	}

	return model.New(map[string]interface{}{
		"ID":        session.ID,
		"User":      session.User,
		"Ip":        session.Ip,
		"Agent":     session.Agent,
		"CreatedAt": session.CreatedAt,
		"ExpiresAt": session.ExpiresAt,
	}), nil
}

func (repository session) Get(ctx context.Context, req model.Value) (model.Value, error) {
	session, err := repository.store.GetSession(ctx, req.Get("ID").(uuid.UUID))
	if err != nil {
		return nil, err
	}

	return model.New(map[string]interface{}{
		"ID":        session.ID,
		"User":      session.User,
		"Ip":        session.Ip,
		"Agent":     session.Agent,
		"CreatedAt": session.CreatedAt,
		"ExpiresAt": session.ExpiresAt,
	}), nil
}

func (repository session) Find(ctx context.Context, req model.Value) ([]model.Value, error) {
	res, err := repository.store.FindSession(ctx, req.Get("User").(uuid.UUID))
	if err != nil {
		return nil, err
	}

	var sessions []model.Value

	for _, session := range res {
		sessions = append(sessions, model.New(map[string]interface{}{
			"ID":        session.ID,
			"User":      session.User,
			"Ip":        session.Ip,
			"Agent":     session.Agent,
			"CreatedAt": session.CreatedAt,
			"ExpiresAt": session.ExpiresAt,
		}))
	}
	return sessions, nil
}

func (repository session) List(ctx context.Context, req model.Value) ([]model.Value, error) {
	_, limit := req.Get("Limit").(int32)
	_, offset := req.Get("Offset").(int32)
	if limit && offset {
		res, err := repository.store.ListSessionPage(ctx, &db.ListSessionPageParams{
			Limit:  req.Get("Limit").(int32),
			Offset: req.Get("Offset").(int32),
		})
		if err != nil {
			return nil, err
		}

		var sessions []model.Value

		for _, session := range res {
			sessions = append(sessions, model.New(map[string]interface{}{
				"ID":        session.ID,
				"User":      session.User,
				"Ip":        session.Ip,
				"Agent":     session.Agent,
				"CreatedAt": session.CreatedAt,
				"ExpiresAt": session.ExpiresAt,
			}))
		}

		return sessions, nil
	}

	res, err := repository.store.ListSession(ctx)
	if err != nil {
		return nil, err
	}

	var sessions []model.Value

	for _, session := range res {
		sessions = append(sessions, model.New(map[string]interface{}{
			"ID":        session.ID,
			"User":      session.User,
			"Ip":        session.Ip,
			"Agent":     session.Agent,
			"CreatedAt": session.CreatedAt,
			"ExpiresAt": session.ExpiresAt,
		}))
	}

	return sessions, nil
}

func (repository session) Update(ctx context.Context, req model.Value) (model.Value, error) {
	return nil, errors.New("not implemented method")
}

func (repository session) Delete(ctx context.Context, req model.Value) (model.Value, error) {
	session, err := repository.store.DeleteSession(ctx, req.Get("ID").(uuid.UUID))
	if err != nil {
		return nil, err
	}

	return model.New(map[string]interface{}{
		"ID":        session.ID,
		"User":      session.User,
		"Ip":        session.Ip,
		"Agent":     session.Agent,
		"CreatedAt": session.CreatedAt,
		"ExpiresAt": session.ExpiresAt,
	}), nil
}
