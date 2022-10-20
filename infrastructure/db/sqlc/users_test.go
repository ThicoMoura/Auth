package db

import (
	"context"
	"database/sql"
	"testing"

	"github.com/ThicoMoura/Auth/util"
	"github.com/stretchr/testify/require"
)

func createRandomUser(t *testing.T) User {
	arg := CreateUsersParams{
		Cpf:  util.RandomString(11),
		Name: util.RandomString(10),
		Pass: util.RandomString(10),
	}

	user, err := testQueries.CreateUsers(context.Background(), arg)

	require.NoError(t, err)
	require.NotEmpty(t, user)

	require.Equal(t, arg.Cpf, user.Cpf)
	require.Equal(t, arg.Name, user.Name)
	require.Equal(t, arg.Pass, user.Pass)
	require.Equal(t, true, user.Active)

	require.NotZero(t, user.ID)

	return user
}

func TestCreateUsers(t *testing.T) {
	createRandomUser(t)
}

func TestGetUser(t *testing.T) {
	user1 := createRandomUser(t)
	user2, err := testQueries.GetUsers(context.Background(), user1.ID)

	require.NoError(t, err)
	require.NotEmpty(t, user2)

	require.Equal(t, user1.Cpf, user2.Cpf)
	require.Equal(t, user1.Name, user2.Name)
	require.Equal(t, user1.Pass, user2.Pass)
	require.Equal(t, user1.Active, user2.Active)
}

func TestUpdateUser(t *testing.T) {
	user1 := createRandomUser(t)
	arg := UpdateUsersParams{
		ID:     user1.ID,
		Cpf:    user1.Cpf,
		Name:   util.RandomString(10),
		Pass:   user1.Pass,
		Active: user1.Active,
	}

	user2, err := testQueries.UpdateUsers(context.Background(), arg)

	require.NoError(t, err)
	require.NotEmpty(t, user2)

	require.Equal(t, user1.ID, user2.ID)
	require.Equal(t, user1.Cpf, user2.Cpf)
	require.Equal(t, arg.Name, user2.Name)
	require.Equal(t, user1.Pass, user2.Pass)
	require.Equal(t, user1.Active, user2.Active)
}

func TestDeleteUser(t *testing.T) {
	user1 := createRandomUser(t)
	err := testQueries.DeleteUsers(context.Background(), user1.ID)
	require.NoError(t, err)

	user2, err := testQueries.GetUsers(context.Background(), user1.ID)
	require.Error(t, err)
	require.EqualError(t, err, sql.ErrNoRows.Error())
	require.Empty(t, user2)
}

func TestListUser(t *testing.T) {
	for i := 0; i < 10; i++ {
		createRandomUser(t)
	}

	arg := ListUsersParams{
		Limit:  5,
		Offset: 5,
	}

	list, err := testQueries.ListUsers(context.Background(), arg)

	require.NoError(t, err)
	require.Len(t, list, 5)

	for _, user := range list {
		require.NotEmpty(t, user)
	}
}
