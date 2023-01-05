package repository_test

import (
	"context"
	"testing"

	"github.com/ThicoMoura/Auth/db/model"
	"github.com/ThicoMoura/Auth/util"
	"github.com/google/uuid"
	"github.com/stretchr/testify/require"
)

func NewGroup(t *testing.T) model.Value {
	group, err := repo.Table("group").New(context.Background(), model.New(map[string]interface{}{
		"Name": util.RandomString(10),
	}))

	require.NoError(t, err)
	require.NotEmpty(t, group)

	return group
}

func DeleteGroup(t *testing.T, ID uuid.UUID) model.Value {
	group, err := repo.Table("group").Delete(context.Background(), model.New(map[string]interface{}{
		"ID": ID,
	}))

	require.NoError(t, err)
	require.NotEmpty(t, group)

	return group
}

func TestNewGroup(t *testing.T) {
	DeleteGroup(t, NewGroup(t).Get("ID").(uuid.UUID))
}

func TestGetGroup(t *testing.T) {
	group := NewGroup(t)

	res, err := repo.Table("group").Get(context.Background(), model.New(map[string]interface{}{
		"ID": group.Get("ID"),
	}))

	require.NoError(t, err)
	require.NotEmpty(t, res)
	require.Equal(t, group, res)

	DeleteGroup(t, group.Get("ID").(uuid.UUID))
}

func TestFindGroup(t *testing.T) {
	var groups []model.Value
	name := 0

	for i := 0; i < 10; i++ {
		group := NewGroup(t)

		if group.Get("Name").(string)[0:1] == "a" {
			name++
		}

		groups = append(groups, group)
	}

	res, err := repo.Table("group").Find(context.Background(), model.New(map[string]interface{}{
		"Name": "a%",
	}))

	require.NoError(t, err)
	require.Len(t, res, name)

	for _, group := range res {
		require.NotEmpty(t, group)
	}

	for _, group := range groups {
		DeleteGroup(t, group.Get("ID").(uuid.UUID))
	}
}

func TestListGroup(t *testing.T) {
	for i := 0; i < 10; i++ {
		NewGroup(t)
	}

	groups, err := repo.Table("group").List(context.Background(), model.New(map[string]interface{}{
		"Limit":  int32(5),
		"Offset": int32(5),
	}))

	require.NoError(t, err)
	require.Len(t, groups, 5)

	for _, group := range groups {
		require.NotEmpty(t, group)
	}

	groups, err = repo.Table("group").List(context.Background(), model.New(nil))

	require.NoError(t, err)
	require.Len(t, groups, 10)

	for _, group := range groups {
		require.NotEmpty(t, group)

		DeleteGroup(t, group.Get("ID").(uuid.UUID))
	}
}

func TestUpdateGroup(t *testing.T) {
	group := NewGroup(t)

	arg := model.New(map[string]interface{}{
		"ID":     group.Get("ID"),
		"Active": util.RandomBool(),
	})

	updated, err := repo.Table("group").Update(context.Background(), arg)

	require.NoError(t, err)
	require.NotEmpty(t, group)

	require.Equal(t, group.Get("ID"), updated.Get("ID"))
	require.Equal(t, group.Get("Name"), updated.Get("Name"))
	require.Equal(t, arg.Get("Tables"), updated.Get("Tables"))
	require.Equal(t, arg.Get("Active"), updated.Get("Active"))

	DeleteGroup(t, group.Get("ID").(uuid.UUID))
}
