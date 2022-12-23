package db_test

import (
	"context"
	"testing"

	db "github.com/ThicoMoura/Auth/db/sqlc"
	"github.com/ThicoMoura/Auth/util"
	"github.com/google/uuid"
	"github.com/stretchr/testify/require"
)

func NewUser(t *testing.T) *db.User {
	arg := &db.NewUserParams{
		Cpf:  util.RandomString(11),
		Name: util.RandomString(10),
		Pass: util.RandomString(10),
	}

	user, err := testQueries.NewUser(context.Background(), dbtx, arg)

	require.NoError(t, err)
	require.NotEmpty(t, user)

	require.Equal(t, arg.Cpf, user.Cpf)
	require.Equal(t, arg.Name, user.Name)
	require.Equal(t, arg.Pass, user.Pass)
	require.Equal(t, true, user.Active)

	require.NotZero(t, user.ID)

	return user
}

func DeleteUser(t *testing.T, ID uuid.UUID) *db.User {
	user, err := testQueries.DeleteUser(context.Background(), dbtx, ID)
	require.NoError(t, err)
	require.NotEmpty(t, user)
	require.Equal(t, ID, user.ID)

	return user
}

func TestCreateUsers(t *testing.T) {
	DeleteUser(t, NewUser(t).ID)
}

// func TestGetUser(t *testing.T) {
// 	user1 := createRandomUser(t)
// 	user2, err := testQueries.GetUser(context.Background(), db, user1.ID)

// 	require.NoError(t, err)
// 	require.NotEmpty(t, user2)

// 	require.Equal(t, user1.Cpf, user2.Cpf)
// 	require.Equal(t, user1.Name, user2.Name)
// 	require.Equal(t, user1.Pass, user2.Pass)
// 	require.Equal(t, user1.Active, user2.Active)
// }

// func TestGetUserByCPF(t *testing.T) {
// 	user1 := createRandomUser(t)
// 	user2, err := testQueries.GetUserByCPF(context.Background(), db, user1.Cpf)

// 	require.NoError(t, err)
// 	require.NotEmpty(t, user2)

// 	require.Equal(t, user1.Cpf, user2.Cpf)
// 	require.Equal(t, user1.Name, user2.Name)
// 	require.Equal(t, user1.Pass, user2.Pass)
// 	require.Equal(t, user1.Active, user2.Active)
// }

// func TestUpdateUserName(t *testing.T) {
// 	user1 := createRandomUser(t)
// 	arg := &UpdateUserParams{
// 		ID:   user1.ID,
// 		Name: util.RandomString(10),
// 	}

// 	user2, err := testQueries.UpdateUser(context.Background(), db, arg)

// 	require.NoError(t, err)
// 	require.NotEmpty(t, user2)

// 	require.Equal(t, user1.ID, user2.ID)
// 	require.Equal(t, user1.Cpf, user2.Cpf)
// 	require.Equal(t, arg.Name, user2.Name)
// 	require.Equal(t, user1.Pass, user2.Pass)
// 	require.Equal(t, user1.Active, user2.Active)
// }

// func TestUpdateUserPass(t *testing.T) {
// 	user1 := createRandomUser(t)
// 	arg := &UpdateUserPassParams{
// 		ID:   user1.ID,
// 		Pass: util.RandomString(10),
// 	}

// 	user2, err := testQueries.UpdateUserPass(context.Background(), db, arg)

// 	require.NoError(t, err)
// 	require.NotEmpty(t, user2)

// 	require.Equal(t, user1.ID, user2.ID)
// 	require.Equal(t, user1.Cpf, user2.Cpf)
// 	require.Equal(t, user1.Name, user2.Name)
// 	require.Equal(t, arg.Pass, user2.Pass)
// 	require.Equal(t, user1.Active, user2.Active)
// }

// func TestListUser(t *testing.T) {
// 	for i := 0; i < 10; i++ {
// 		createRandomUser(t)
// 	}

// 	arg := &ListUserPageParams{
// 		Limit:  5,
// 		Offset: 5,
// 	}

// 	list, err := testQueries.ListUserPage(context.Background(), db, arg)

// 	require.NoError(t, err)
// 	require.Len(t, list, 5)

// 	for _, user := range list {
// 		require.NotEmpty(t, user)
// 	}
// }
