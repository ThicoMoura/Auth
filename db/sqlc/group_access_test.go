package db_test

import (
	"context"
	"testing"

	db "github.com/ThicoMoura/Auth/db/sqlc"
	"github.com/google/uuid"
	"github.com/stretchr/testify/require"
)

func NewGroupAccess(t *testing.T, group, access uuid.UUID) *db.GroupAccess {
	groupAccess, err := testQueries.NewGroupAccess(context.Background(), &db.NewGroupAccessParams{
		GroupID:  group,
		AccessID: access,
	})

	require.NoError(t, err)
	require.NotEmpty(t, groupAccess)

	require.Equal(t, group, groupAccess.GroupID)
	require.Equal(t, access, groupAccess.AccessID)

	return groupAccess
}

func DeleteGroupAccess(t *testing.T, group, access uuid.UUID) *db.GroupAccess {
	groupAccess, err := testQueries.DeleteGroupAccess(context.Background(), &db.DeleteGroupAccessParams{
		GroupID:  group,
		AccessID: access,
	})

	require.NoError(t, err)
	require.NotEmpty(t, groupAccess)

	require.Equal(t, group, groupAccess.GroupID)
	require.Equal(t, access, groupAccess.AccessID)

	return groupAccess
}

func TestNewGroupAccess(t *testing.T) {
	system := NewSystem(t).ID
	groupAccess := NewGroupAccess(t, NewGroup(t).ID, NewAccess(t, system).ID)
	DeleteGroupAccess(t, groupAccess.GroupID, groupAccess.AccessID)
	DeleteGroup(t, groupAccess.GroupID)
	DeleteAccess(t, groupAccess.AccessID)
	DeleteSystem(t, system)
}

func TestFindGroupAccess(t *testing.T) {
	system := NewSystem(t).ID
	group := NewGroup(t).ID

	for i := 0; i < 10; i++ {
		NewGroupAccess(t, group, NewAccess(t, system).ID)
	}

	res, err := testQueries.FindGroupAccess(context.Background(), &db.FindGroupAccessParams{
		GroupID: group,
	})

	require.NoError(t, err)
	require.Len(t, res, 10)

	for _, groupAccess := range res {
		require.NotEmpty(t, groupAccess)
		DeleteGroupAccess(t, groupAccess.GroupID, groupAccess.AccessID)
		DeleteAccess(t, groupAccess.AccessID)
	}

	DeleteGroup(t, group)

	access := NewAccess(t, system).ID

	for i := 0; i < 10; i++ {
		NewGroupAccess(t, NewGroup(t).ID, access)
	}

	res, err = testQueries.FindGroupAccess(context.Background(), &db.FindGroupAccessParams{
		AccessID: access,
	})

	require.NoError(t, err)
	require.Len(t, res, 10)

	for _, groupAccess := range res {
		require.NotEmpty(t, groupAccess)
		DeleteGroupAccess(t, groupAccess.GroupID, groupAccess.AccessID)
		DeleteGroup(t, groupAccess.GroupID)
	}

	DeleteAccess(t, access)

	DeleteSystem(t, system)
}
