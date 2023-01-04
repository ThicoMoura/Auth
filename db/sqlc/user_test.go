package db_test

import (
	"context"
	"testing"

	db "github.com/ThicoMoura/Auth/db/sqlc"
	"github.com/ThicoMoura/Auth/util"
	"github.com/google/uuid"
	"github.com/stretchr/testify/require"
)

func NewUser(t *testing.T, ID uuid.UUID) *db.User {
	arg := &db.NewUserParams{
		Group: ID,
		Email: util.RandomString(11),
		Name:  util.RandomString(10),
		Pass:  util.RandomString(10),
	}

	user, err := testQueries.NewUser(context.Background(), arg)

	require.NoError(t, err)
	require.NotEmpty(t, user)

	require.NotEmpty(t, user.ID)
	require.Equal(t, arg.Group, user.Group)
	require.Equal(t, arg.Email, user.Email)
	require.Equal(t, arg.Name, user.Name)
	require.Equal(t, arg.Pass, user.Pass)
	require.Equal(t, true, user.Active)

	return user
}

func DeleteUser(t *testing.T, ID uuid.UUID) *db.User {
	user, err := testQueries.DeleteUser(context.Background(), ID)
	require.NoError(t, err)
	require.NotEmpty(t, user)
	require.Equal(t, ID, user.ID)

	return user
}

func TestNewUsers(t *testing.T) {
	user := NewUser(t, NewGroup(t).ID)
	DeleteUser(t, user.ID)
	DeleteGroup(t, user.Group)
}

func TestGetUser(t *testing.T) {
	user := NewUser(t, NewGroup(t).ID)

	res, err := testQueries.GetUser(context.Background(), &db.GetUserParams{
		ID: user.ID,
	})

	require.NoError(t, err)
	require.NotEmpty(t, res)
	require.Equal(t, user, res)

	res, err = testQueries.GetUser(context.Background(), &db.GetUserParams{
		Email: user.Email,
	})

	require.NoError(t, err)
	require.NotEmpty(t, res)
	require.Equal(t, user, res)

	DeleteUser(t, user.ID)
	DeleteGroup(t, user.Group)
}

func TestFindUser(t *testing.T) {
	group := NewGroup(t)

	name := 0
	for i := 0; i < 10; i++ {
		user := NewUser(t, group.ID)
		if user.Name[0:1] == "a" {
			name++
		}
	}

	users, err := testQueries.FindUser(context.Background(), &db.FindUserParams{
		Name: "a%",
	})

	require.NoError(t, err)
	require.Len(t, users, name)

	for _, user := range users {
		require.NotEmpty(t, user)
	}

	users, err = testQueries.FindUser(context.Background(), &db.FindUserParams{
		Group: group.ID,
		Name:  "a%",
	})

	require.NoError(t, err)
	require.Len(t, users, 10)

	for _, user := range users {
		require.NotEmpty(t, user)
	}

	users, err = testQueries.FindUser(context.Background(), &db.FindUserParams{
		Group: group.ID,
	})

	require.NoError(t, err)
	require.Len(t, users, 10)

	for _, user := range users {
		require.NotEmpty(t, user)
		DeleteUser(t, user.ID)
	}

	DeleteGroup(t, group.ID)
}

func TestListUser(t *testing.T) {
	for i := 0; i < 10; i++ {
		NewUser(t, NewGroup(t).ID)
	}

	list, err := testQueries.ListUserPage(context.Background(), &db.ListUserPageParams{
		Limit:  5,
		Offset: 5,
	})

	require.NoError(t, err)
	require.Len(t, list, 5)

	for _, user := range list {
		require.NotEmpty(t, user)
	}

	list, err = testQueries.ListUser(context.Background())

	require.NoError(t, err)
	require.Len(t, list, 10)

	for _, user := range list {
		require.NotEmpty(t, user)
		DeleteUser(t, user.ID)
		DeleteGroup(t, user.Group)
	}
}

func TestUpdateUser(t *testing.T) {
	user := NewUser(t, NewGroup(t).ID)
	arg := &db.UpdateUserParams{
		ID:   user.ID,
		Name: util.RandomString(10),
		Pass: util.RandomString(10),
	}

	res, err := testQueries.UpdateUser(context.Background(), arg)

	require.NoError(t, err)
	require.NotEmpty(t, res)

	require.Equal(t, user.ID, res.ID)
	require.Equal(t, user.Email, res.Email)
	require.Equal(t, arg.Name, res.Name)
	require.Equal(t, arg.Pass, res.Pass)
	require.Equal(t, user.Active, res.Active)

	argActive := &db.UpdateUserActiveParams{
		ID:     user.ID,
		Active: util.RandomBool(),
	}

	res, err = testQueries.UpdateUserActive(context.Background(), argActive)

	require.NoError(t, err)
	require.NotEmpty(t, res)

	require.Equal(t, user.ID, res.ID)
	require.Equal(t, user.Email, res.Email)
	require.Equal(t, arg.Name, res.Name)
	require.Equal(t, arg.Pass, res.Pass)
	require.Equal(t, argActive.Active, res.Active)

	DeleteUser(t, user.ID)
	DeleteGroup(t, user.Group)
}
