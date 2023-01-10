package repository_test

import (
	"context"
	"testing"

	"github.com/ThicoMoura/Auth/db/model"
	"github.com/ThicoMoura/Auth/util"
	"github.com/google/uuid"
	"github.com/stretchr/testify/require"
)

func NewSystem(t *testing.T) model.Value {
	system, err := repo.Table("system").New(context.Background(), model.New(map[string]interface{}{
		"Name":   util.RandomString(10),
		"Tables": []string{},
	}))

	require.NoError(t, err)
	require.NotEmpty(t, system)

	return system
}

func DeleteSystem(t *testing.T, ID uuid.UUID) model.Value {
	system, err := repo.Table("system").Delete(context.Background(), model.New(map[string]interface{}{
		"ID": ID,
	}))

	require.NoError(t, err)
	require.NotEmpty(t, system)

	return system
}

func TestNewSystem(t *testing.T) {
	DeleteSystem(t, NewSystem(t).Get("ID").(uuid.UUID))
}

func TestGetSystem(t *testing.T) {
	system := NewSystem(t)

	res, err := repo.Table("system").Get(context.Background(), model.New(map[string]interface{}{
		"ID": system.Get("ID"),
	}))

	require.NoError(t, err)
	require.NotEmpty(t, res)
	require.Equal(t, system, res)

	DeleteSystem(t, system.Get("ID").(uuid.UUID))
}

func TestFindSystem(t *testing.T) {
	var systems []model.Value
	name := 0

	for i := 0; i < 10; i++ {
		system := NewSystem(t)

		if system.Get("Name").(string)[0:1] == "a" {
			name++
		}

		systems = append(systems, system)
	}

	res, err := repo.Table("system").Find(context.Background(), model.New(map[string]interface{}{
		"Name": "a%",
	}))

	require.NoError(t, err)
	require.Len(t, res, name)

	for _, system := range res {
		require.NotEmpty(t, system)
	}

	for _, system := range systems {
		DeleteSystem(t, system.Get("ID").(uuid.UUID))
	}
}

func TestListSystem(t *testing.T) {
	for i := 0; i < 10; i++ {
		NewSystem(t)
	}

	systems, err := repo.Table("system").List(context.Background(), model.New(map[string]interface{}{
		"Limit":  int32(5),
		"Offset": int32(5),
	}))

	require.NoError(t, err)
	require.Len(t, systems, 5)

	for _, system := range systems {
		require.NotEmpty(t, system)
	}

	systems, err = repo.Table("system").List(context.Background(), model.New(nil))

	require.NoError(t, err)
	require.Len(t, systems, 10)

	for _, system := range systems {
		require.NotEmpty(t, system)

		DeleteSystem(t, system.Get("ID").(uuid.UUID))
	}
}

func TestUpdateSystem(t *testing.T) {
	system := NewSystem(t)
	defer DeleteSystem(t, system.Get("ID").(uuid.UUID))

	arg := model.New(map[string]interface{}{
		"ID":     system.Get("ID"),
		"Active": util.RandomBool(),
	})

	updated, err := repo.Table("system").Update(context.Background(), arg)

	require.NoError(t, err)
	require.NotEmpty(t, system)

	require.Equal(t, system.Get("ID"), updated.Get("ID"))
	require.Equal(t, system.Get("Name"), updated.Get("Name"))
	require.Equal(t, arg.Get("Tables"), updated.Get("Tables"))
	require.Equal(t, arg.Get("Active"), updated.Get("Active"))
}
