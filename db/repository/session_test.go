package repository_test

import (
	"context"
	"testing"
	"time"

	"github.com/ThicoMoura/Auth/db/model"
	"github.com/ThicoMoura/Auth/util"
	"github.com/google/uuid"
	"github.com/stretchr/testify/require"
)

func NewSession(t *testing.T, ID uuid.UUID) model.Value {
	session, err := repo.Table("session").New(context.Background(), model.New(map[string]interface{}{
		"ID":        uuid.New(),
		"User":      ID,
		"Token":     util.RandomString(10),
		"Ip":        util.RandomString(10),
		"Agent":     util.RandomString(10),
		"ExpiresAt": time.Now(),
	}))

	require.NoError(t, err)
	require.NotEmpty(t, session)

	return session
}

func DeleteSession(t *testing.T, ID uuid.UUID) model.Value {
	session, err := repo.Table("session").Delete(context.Background(), model.New(map[string]interface{}{
		"ID": ID,
	}))

	require.NoError(t, err)
	require.NotEmpty(t, session)

	return session
}

func TestNewSession(t *testing.T) {
	group := NewGroup(t)
	session := NewSession(t, NewUser(t, group.Get("ID").(uuid.UUID)).Get("ID").(uuid.UUID))
	DeleteSession(t, session.Get("ID").(uuid.UUID))
	DeleteUser(t, session.Get("User").(uuid.UUID))
	DeleteGroup(t, group.Get("ID").(uuid.UUID))
}

func TestGetSession(t *testing.T) {
	group := NewGroup(t)
	session := NewSession(t, NewUser(t, group.Get("ID").(uuid.UUID)).Get("ID").(uuid.UUID))

	res, err := repo.Table("session").Get(context.Background(), model.New(map[string]interface{}{
		"ID": session.Get("ID"),
	}))

	require.NoError(t, err)
	require.NotEmpty(t, res)
	require.Equal(t, session, res)

	DeleteSession(t, session.Get("ID").(uuid.UUID))
	DeleteUser(t, session.Get("User").(uuid.UUID))
	DeleteGroup(t, group.Get("ID").(uuid.UUID))
}

func TestFindSession(t *testing.T) {
	user := NewUser(t, NewGroup(t).Get("ID").(uuid.UUID))

	for i := 0; i < 10; i++ {
		NewSession(t, user.Get("ID").(uuid.UUID))
	}

	sessions, err := repo.Table("session").Find(context.Background(), model.New(map[string]interface{}{
		"User": user.Get("ID").(uuid.UUID),
	}))

	require.NoError(t, err)
	require.Len(t, sessions, 10)

	for _, session := range sessions {
		require.NotEmpty(t, session)
		DeleteSession(t, session.Get("ID").(uuid.UUID))
	}

	DeleteUser(t, user.Get("ID").(uuid.UUID))
	DeleteGroup(t, user.Get("Group").(uuid.UUID))
}

func TestListSession(t *testing.T) {
	user := NewUser(t, NewGroup(t).Get("ID").(uuid.UUID))

	for i := 0; i < 10; i++ {
		NewSession(t, user.Get("ID").(uuid.UUID))
	}

	sessions, err := repo.Table("session").List(context.Background(), model.New(map[string]interface{}{
		"Limit":  int32(5),
		"Offset": int32(5),
	}))

	require.NoError(t, err)
	require.Len(t, sessions, 5)

	for _, session := range sessions {
		require.NotEmpty(t, session)
	}

	sessions, err = repo.Table("session").List(context.Background(), model.New(nil))

	require.NoError(t, err)
	require.Len(t, sessions, 10)

	for _, session := range sessions {
		require.NotEmpty(t, session)
		DeleteSession(t, session.Get("ID").(uuid.UUID))
	}

	DeleteUser(t, user.Get("ID").(uuid.UUID))
	DeleteGroup(t, user.Get("Group").(uuid.UUID))
}
