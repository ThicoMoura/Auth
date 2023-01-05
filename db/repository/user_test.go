package repository_test

import (
	"context"
	"testing"

	"github.com/ThicoMoura/Auth/db/model"
	"github.com/ThicoMoura/Auth/util"
	"github.com/google/uuid"
	"github.com/stretchr/testify/require"
)

func NewUser(t *testing.T, ID uuid.UUID) model.Value {
	user, err := repo.Table("user").New(context.Background(), model.New(map[string]interface{}{
		"Group": ID,
		"Email": util.RandomString(10),
		"Name":  util.RandomString(10),
		"Pass":  util.RandomString(10),
	}))

	require.NoError(t, err)
	require.NotEmpty(t, user)

	return user
}

func DeleteUser(t *testing.T, ID uuid.UUID) model.Value {
	user, err := repo.Table("user").Delete(context.Background(), model.New(map[string]interface{}{
		"ID": ID,
	}))

	require.NoError(t, err)
	require.NotEmpty(t, user)

	return user
}

func TestNewUser(t *testing.T) {
	user := NewUser(t, NewGroup(t).Get("ID").(uuid.UUID))
	DeleteUser(t, user.Get("ID").(uuid.UUID))
	DeleteGroup(t, user.Get("Group").(uuid.UUID))
}

func TestGetUser(t *testing.T) {
	user := NewUser(t, NewGroup(t).Get("ID").(uuid.UUID))

	res, err := repo.Table("user").Get(context.Background(), model.New(map[string]interface{}{
		"ID": user.Get("ID").(uuid.UUID),
	}))

	require.NoError(t, err)
	require.NotEmpty(t, res)
	require.Equal(t, user, res)

	res, err = repo.Table("user").Get(context.Background(), model.New(map[string]interface{}{
		"Email": user.Get("Email").(string),
	}))

	require.NoError(t, err)
	require.NotEmpty(t, res)
	require.Equal(t, user, res)

	DeleteUser(t, user.Get("ID").(uuid.UUID))
	DeleteGroup(t, user.Get("Group").(uuid.UUID))
}

func TestFindUser(t *testing.T) {
	var users []model.Value

	group := NewGroup(t)
	name := 0
	for i := 0; i < 10; i++ {
		user := NewUser(t, group.Get("ID").(uuid.UUID))

		if user.Get("Name").(string)[0:1] == "a" {
			name++
		}

		users = append(users, user)
	}

	res, err := repo.Table("user").Find(context.Background(), model.New(map[string]interface{}{
		"Group": group.Get("ID").(uuid.UUID),
	}))

	require.NoError(t, err)
	require.Len(t, res, 10)

	for _, user := range res {
		require.NotEmpty(t, user)
	}

	res, err = repo.Table("user").Find(context.Background(), model.New(map[string]interface{}{
		"Name": "a%",
	}))

	require.NoError(t, err)
	require.Len(t, res, name)

	for _, user := range res {
		require.NotEmpty(t, user)
	}

	res, err = repo.Table("user").Find(context.Background(), model.New(map[string]interface{}{
		"Group": group.Get("ID").(uuid.UUID),
		"Name":  "a%",
	}))

	require.NoError(t, err)
	require.Len(t, res, 10)

	for _, user := range res {
		require.NotEmpty(t, user)
	}

	res, err = repo.Table("user").Find(context.Background(), model.New(map[string]interface{}{
		"Group":  group.Get("ID").(uuid.UUID),
		"Limit":  int32(5),
		"Offset": int32(5),
	}))

	require.NoError(t, err)
	require.Len(t, res, 5)

	for _, user := range res {
		require.NotEmpty(t, user)
	}

	for _, user := range users {
		DeleteUser(t, user.Get("ID").(uuid.UUID))
	}

	DeleteGroup(t, group.Get("ID").(uuid.UUID))
}

func TestListUser(t *testing.T) {
	for i := 0; i < 10; i++ {
		NewUser(t, NewGroup(t).Get("ID").(uuid.UUID))
	}

	users, err := repo.Table("user").List(context.Background(), model.New(map[string]interface{}{
		"Limit":  int32(5),
		"Offset": int32(5),
	}))

	require.NoError(t, err)
	require.Len(t, users, 5)

	for _, user := range users {
		require.NotEmpty(t, user)
	}

	users, err = repo.Table("user").List(context.Background(), model.New(nil))

	require.NoError(t, err)
	require.Len(t, users, 10)

	for _, user := range users {
		require.NotEmpty(t, user)
		DeleteUser(t, user.Get("ID").(uuid.UUID))
		DeleteGroup(t, user.Get("Group").(uuid.UUID))
	}
}

func TestUpdateUser(t *testing.T) {
	user := NewUser(t, NewGroup(t).Get("ID").(uuid.UUID))

	arg := model.New(map[string]interface{}{
		"ID":     user.Get("ID"),
		"Name":   util.RandomString(10),
		"Pass":   util.RandomString(10),
		"Active": util.RandomBool(),
	})

	res, err := repo.Table("user").Update(context.Background(), arg)

	require.NoError(t, err)
	require.NotEmpty(t, res)

	require.Equal(t, user.Get("ID"), res.Get("ID"))
	require.Equal(t, user.Get("Email"), res.Get("Email"))
	require.Equal(t, arg.Get("Name"), res.Get("Name"))
	require.Equal(t, arg.Get("Pass"), res.Get("Pass"))
	require.Equal(t, arg.Get("Active"), res.Get("Active"))

	DeleteUser(t, user.Get("ID").(uuid.UUID))
	DeleteGroup(t, user.Get("Group").(uuid.UUID))
}
