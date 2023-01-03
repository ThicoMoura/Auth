package db_test

import (
	"context"
	"testing"

	db "github.com/ThicoMoura/Auth/db/sqlc"
	"github.com/ThicoMoura/Auth/util"
	"github.com/google/uuid"
	"github.com/stretchr/testify/require"
)

func NewAccess(t *testing.T, ID uuid.UUID) *db.Access {
	arg := &db.NewAccessParams{
		System: ID,
		Table:  util.RandomString(10),
	}

	for i := 0; i < 10; i++ {
		arg.Type = append(arg.Type, util.RandomString(10))
	}

	access, err := testQueries.NewAccess(context.Background(), arg)

	require.NoError(t, err)
	require.NotEmpty(t, access)

	require.NotEmpty(t, access.ID)
	require.Equal(t, arg.System, access.System)
	require.Equal(t, arg.Table, access.Table)
	require.Equal(t, arg.Type, access.Type)

	return access
}

func DeleteAccess(t *testing.T, ID uuid.UUID) *db.Access {
	access, err := testQueries.DeleteAccess(context.Background(), ID)
	require.NoError(t, err)
	require.NotEmpty(t, access)
	require.Equal(t, ID, access.ID)

	return access
}

func TestNewAccesss(t *testing.T) {
	access := NewAccess(t, NewSystem(t).ID)
	DeleteAccess(t, access.ID)
	DeleteSystem(t, access.System)
}

func TestGetAccess(t *testing.T) {
	access := NewAccess(t, NewSystem(t).ID)

	res, err := testQueries.GetAccess(context.Background(), access.ID)

	require.NoError(t, err)
	require.NotEmpty(t, res)
	require.Equal(t, access, res)

	DeleteAccess(t, access.ID)
	DeleteSystem(t, access.System)
}

func TestFindAccess(t *testing.T) {
	system := NewSystem(t)

	table := 0
	for i := 0; i < 10; i++ {
		access := NewAccess(t, system.ID)
		if access.Table[0:1] == "a" {
			table++
		}
	}

	accesss, err := testQueries.FindAccess(context.Background(), &db.FindAccessParams{
		Table: "a%",
	})

	require.NoError(t, err)
	require.Len(t, accesss, table)

	for _, access := range accesss {
		require.NotEmpty(t, access)
	}

	accesss, err = testQueries.FindAccess(context.Background(), &db.FindAccessParams{
		System: system.ID,
		Table:  "a%",
	})

	require.NoError(t, err)
	require.Len(t, accesss, 10)

	for _, access := range accesss {
		require.NotEmpty(t, access)
	}

	accesss, err = testQueries.FindAccess(context.Background(), &db.FindAccessParams{
		System: system.ID,
	})

	require.NoError(t, err)
	require.Len(t, accesss, 10)

	for _, access := range accesss {
		require.NotEmpty(t, access)
		DeleteAccess(t, access.ID)
	}

	DeleteSystem(t, system.ID)
}

func TestListAccess(t *testing.T) {
	for i := 0; i < 10; i++ {
		NewAccess(t, NewSystem(t).ID)
	}

	list, err := testQueries.ListAccessPage(context.Background(), &db.ListAccessPageParams{
		Limit:  5,
		Offset: 5,
	})

	require.NoError(t, err)
	require.Len(t, list, 5)

	for _, access := range list {
		require.NotEmpty(t, access)
	}

	list, err = testQueries.ListAccess(context.Background())

	require.NoError(t, err)
	require.Len(t, list, 10)

	for _, access := range list {
		require.NotEmpty(t, access)
		DeleteAccess(t, access.ID)
		DeleteSystem(t, access.System)
	}
}

func TestUpdateAccess(t *testing.T) {
	access := NewAccess(t, NewSystem(t).ID)
	arg := &db.UpdateAccessParams{
		ID: access.ID,
	}

	for i := 0; i < 10; i++ {
		arg.Type = append(arg.Type, util.RandomString(10))
	}

	res, err := testQueries.UpdateAccess(context.Background(), arg)

	require.NoError(t, err)
	require.NotEmpty(t, res)

	require.Equal(t, access.ID, res.ID)
	require.Equal(t, access.System, res.System)
	require.Equal(t, access.Table, res.Table)
	require.Equal(t, arg.Type, res.Type)

	DeleteAccess(t, access.ID)
	DeleteSystem(t, access.System)
}
