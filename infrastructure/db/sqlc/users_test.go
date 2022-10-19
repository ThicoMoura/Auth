package db

import (
	"context"
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
