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

	require.NotEmpty(t, group.ID)
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
	defer DeleteGroup(t, NewGroup(t).ID)
}

func TestGetGroup(t *testing.T) {
	group := NewGroup(t)

	defer DeleteGroup(t, group.ID)

	res, err := testQueries.GetGroup(context.Background(), group.ID)

	require.NoError(t, err)
	require.NotEmpty(t, res)

	require.Equal(t, group, res)
}

func TestFindGroup(t *testing.T) {
	var list []*db.Group
	name := 0
	for i := 0; i < 10; i++ {
		group := NewGroup(t)
		if group.Name[0:1] == "a" {
			name++
		}
		list = append(list, group)
	}

	defer func() {
		for _, group := range list {
			DeleteGroup(t, group.ID)
		}
	}()

	res, err := testQueries.FindGroup(context.Background(), "a%")

	require.NoError(t, err)
	require.Len(t, res, name)
	for _, group := range res {
		require.NotEmpty(t, group)
	}

	res, err = testQueries.FindGroupPage(context.Background(), &db.FindGroupPageParams{
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

	for _, group := range res {
		require.NotEmpty(t, group)
	}
}

func TestListGroup(t *testing.T) {
	var IDs []uuid.UUID
	for i := 0; i < 10; i++ {
		IDs = append(IDs, NewGroup(t).ID)
	}

	defer func() {
		for _, ID := range IDs {
			DeleteGroup(t, ID)
		}
	}()

	list, err := testQueries.ListGroupPage(context.Background(), &db.ListGroupPageParams{
		Limit:  5,
		Offset: 5,
	})

	require.NoError(t, err)
	require.Len(t, list, 5)

	for _, group := range list {
		require.NotEmpty(t, group)
	}

	list, err = testQueries.ListGroup(context.Background())

	require.NoError(t, err)
	require.Len(t, list, 10)

	for _, group := range list {
		require.NotEmpty(t, group)
	}
}

func TestUpdateGroup(t *testing.T) {
	group := NewGroup(t)

	defer DeleteGroup(t, group.ID)

	arg := db.UpdateGroupParams{
		ID:   group.ID,
		Name: util.RandomString(10),
	}

	update, err := testQueries.UpdateGroup(context.Background(), &arg)

	require.NoError(t, err)
	require.NotEmpty(t, update)

	require.Equal(t, group.ID, update.ID)
	require.Equal(t, arg.Name, update.Name)
	require.Equal(t, group.Active, update.Active)
}
