package db_test

import (
	"context"
	"testing"

	db "github.com/ThicoMoura/Auth/db/sqlc"
	"github.com/google/uuid"
	"github.com/stretchr/testify/require"
)

func NewUserAccess(t *testing.T, user, access uuid.UUID) *db.UserAccess {
	userAccess, err := testQueries.NewUserAccess(context.Background(), &db.NewUserAccessParams{
		UserID:   user,
		AccessID: access,
	})

	require.NoError(t, err)
	require.NotEmpty(t, userAccess)

	require.Equal(t, user, userAccess.UserID)
	require.Equal(t, access, userAccess.AccessID)

	return userAccess
}

func DeleteUserAccess(t *testing.T, user, access uuid.UUID) *db.UserAccess {
	userAccess, err := testQueries.DeleteUserAccess(context.Background(), &db.DeleteUserAccessParams{
		UserID:   user,
		AccessID: access,
	})

	require.NoError(t, err)
	require.NotEmpty(t, userAccess)

	require.Equal(t, user, userAccess.UserID)
	require.Equal(t, access, userAccess.AccessID)

	return userAccess
}

func TestNewUserAccess(t *testing.T) {
	system := NewSystem(t).ID
	group := NewGroup(t).ID
	userAccess := NewUserAccess(t, NewUser(t, group).ID, NewAccess(t, system).ID)
	DeleteUserAccess(t, userAccess.UserID, userAccess.AccessID)
	DeleteUser(t, userAccess.UserID)
	DeleteGroup(t, group)
	DeleteAccess(t, userAccess.AccessID)
	DeleteSystem(t, system)
}

func TestFindUserAccess(t *testing.T) {
	system := NewSystem(t).ID
	group := NewGroup(t).ID
	user := NewUser(t, group).ID

	for i := 0; i < 10; i++ {
		NewUserAccess(t, user, NewAccess(t, system).ID)
	}

	res, err := testQueries.FindUserAccess(context.Background(), &db.FindUserAccessParams{
		UserID: user,
	})

	require.NoError(t, err)
	require.Len(t, res, 10)

	for _, userAccess := range res {
		require.NotEmpty(t, userAccess)
		DeleteUserAccess(t, userAccess.UserID, userAccess.AccessID)
		DeleteAccess(t, userAccess.AccessID)
	}

	DeleteUser(t, user)

	access := NewAccess(t, system).ID

	for i := 0; i < 10; i++ {
		NewUserAccess(t, NewUser(t, group).ID, access)
	}

	res, err = testQueries.FindUserAccess(context.Background(), &db.FindUserAccessParams{
		AccessID: access,
	})

	require.NoError(t, err)
	require.Len(t, res, 10)

	for _, userAccess := range res {
		require.NotEmpty(t, userAccess)
		DeleteUserAccess(t, userAccess.UserID, userAccess.AccessID)
		DeleteUser(t, userAccess.UserID)
	}

	DeleteAccess(t, access)

	DeleteSystem(t, system)
	DeleteGroup(t, group)
}
