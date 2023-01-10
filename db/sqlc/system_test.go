package db_test

import (
	"context"
	"testing"

	db "github.com/ThicoMoura/Auth/db/sqlc"
	"github.com/ThicoMoura/Auth/util"
	"github.com/google/uuid"
	"github.com/stretchr/testify/require"
)

func NewSystem(t *testing.T) *db.System {
	name := util.RandomString(10)

	system, err := testQueries.NewSystem(context.Background(), &db.NewSystemParams{
		Name: name,
	})

	require.NoError(t, err)
	require.NotEmpty(t, system)

	require.NotEmpty(t, system.ID)
	require.Equal(t, name, system.Name)
	require.Equal(t, true, system.Active)

	return system
}

func DeleteSystem(t *testing.T, ID uuid.UUID) *db.System {
	system, err := testQueries.DeleteSystem(context.Background(), ID)

	require.NoError(t, err)
	require.NotEmpty(t, system)

	require.Equal(t, ID, system.ID)

	return nil
}

func TestNewSystem(t *testing.T) {
	DeleteSystem(t, NewSystem(t).ID)
}

func TestGetSystem(t *testing.T) {
	system := NewSystem(t)
	defer DeleteSystem(t, system.ID)

	res, err := testQueries.GetSystem(context.Background(), system.ID)

	require.NoError(t, err)
	require.NotEmpty(t, res)

	require.Equal(t, system, res)
}

func TestFindSystem(t *testing.T) {
	var list []*db.System
	name := 0
	for i := 0; i < 10; i++ {
		system := NewSystem(t)
		if system.Name[0:1] == "a" {
			name++
		}
		list = append(list, system)
	}

	res, err := testQueries.FindSystem(context.Background(), "a%")

	require.NoError(t, err)
	require.Len(t, res, name)
	for _, system := range res {
		require.NotEmpty(t, system)
	}

	res, err = testQueries.FindSystemPage(context.Background(), &db.FindSystemPageParams{
		Name:   "a%",
		Limit:  1,
		Offset: 0,
	})

	result := name

	if result > 1 {
		result = 1
	}

	require.NoError(t, err)
	require.Len(t, res, result)

	for _, system := range res {
		require.NotEmpty(t, system)
	}

	for _, system := range list {
		DeleteSystem(t, system.ID)
	}
}

func TestListSystem(t *testing.T) {
	for i := 0; i < 10; i++ {
		NewSystem(t)
	}

	list, err := testQueries.ListSystemPage(context.Background(), &db.ListSystemPageParams{
		Limit:  5,
		Offset: 5,
	})

	require.NoError(t, err)
	require.Len(t, list, 5)

	for _, system := range list {
		require.NotEmpty(t, system)
	}

	list, err = testQueries.ListSystem(context.Background())

	require.NoError(t, err)
	require.Len(t, list, 10)

	for _, system := range list {
		require.NotEmpty(t, system)
		DeleteSystem(t, system.ID)
	}
}

func TestUpdateSystem(t *testing.T) {
	system := NewSystem(t)
	defer DeleteSystem(t, system.ID)

	arg := db.UpdateSystemParams{
		ID:   system.ID,
		Name: util.RandomString(10),
	}

	update, err := testQueries.UpdateSystem(context.Background(), &arg)

	require.NoError(t, err)
	require.NotEmpty(t, update)

	require.Equal(t, system.ID, update.ID)
	require.Equal(t, arg.Name, update.Name)
	require.Equal(t, system.Active, update.Active)
}
