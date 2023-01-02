package db_test

import (
	"context"
	"testing"

	db "github.com/ThicoMoura/Auth/db/sqlc"
	"github.com/ThicoMoura/Auth/util"
	"github.com/google/uuid"
	"github.com/stretchr/testify/require"
)

func NewGroup(t *testing.T) *db.Group {
	name := util.RandomString(10)

	group, err := testQueries.NewGroup(context.Background(), name)

	require.NoError(t, err)
	require.NotEmpty(t, group)

	require.Equal(t, name, group.Name)
	require.Equal(t, true, group.Active)

	return group
}

func DeleteGroup(t *testing.T, ID uuid.UUID) *db.Group {
	group, err := testQueries.DeleteGroup(context.Background(), ID)

	require.NoError(t, err)
	require.NotEmpty(t, group)

	require.Equal(t, ID, group.ID)

	return nil
}

func TestNewGroup(t *testing.T) {
	DeleteGroup(t, NewGroup(t).ID)
}
