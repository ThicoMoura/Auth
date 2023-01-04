package repository_test

import (
	"context"
	"testing"

	"github.com/ThicoMoura/Auth/db/model"
	"github.com/ThicoMoura/Auth/util"
	"github.com/google/uuid"
	"github.com/stretchr/testify/require"
)

func TestNewSystem(t *testing.T) {
	rSystem := repo.Table("system")

	system, err := rSystem.New(context.Background(), model.New(map[string]interface{}{
		"Name":   util.RandomString(10),
		"Tables": []string{},
	}))

	require.NoError(t, err)
	require.NotEmpty(t, system)

	system, err = rSystem.Delete(context.Background(), model.New(map[string]interface{}{
		"ID": system.Get("ID").(uuid.UUID),
	}))

	require.NoError(t, err)
	require.NotEmpty(t, system)
}

func TestGetSystem(t *testing.T) {
	rSystem := repo.Table("system")

	system, err := rSystem.New(context.Background(), model.New(map[string]interface{}{
		"Name":   util.RandomString(10),
		"Tables": []string{},
	}))

	require.NoError(t, err)
	require.NotEmpty(t, system)

	res, err := rSystem.Get(context.Background(), model.New(map[string]interface{}{
		"ID": system.Get("ID"),
	}))

	require.NoError(t, err)
	require.NotEmpty(t, res)
	require.Equal(t, system, res)

	system, err = rSystem.Delete(context.Background(), model.New(map[string]interface{}{
		"ID": system.Get("ID").(uuid.UUID),
	}))

	require.NoError(t, err)
	require.NotEmpty(t, system)
}

func TestFindSystem(t *testing.T) {
	rSystem := repo.Table("system")

	var systems []model.Value

	name := 0

	for i := 0; i < 10; i++ {
		system, err := rSystem.New(context.Background(), model.New(map[string]interface{}{
			"Name":   util.RandomString(10),
			"Tables": []string{},
		}))

		require.NoError(t, err)
		require.NotEmpty(t, system)

		if system.Get("Name").(string)[0:1] == "a" {
			name++
		}

		systems = append(systems, system)
	}

	res, err := rSystem.Find(context.Background(), model.New(map[string]interface{}{
		"Name": "a%",
	}))

	require.NoError(t, err)
	require.Len(t, res, name)

	for _, system := range res {
		require.NotEmpty(t, system)
	}

	for _, system := range systems {
		res, err := rSystem.Delete(context.Background(), model.New(map[string]interface{}{
			"ID": system.Get("ID").(uuid.UUID),
		}))

		require.NoError(t, err)
		require.NotEmpty(t, res)
	}
}

func TestListSystem(t *testing.T) {
	rSystem := repo.Table("system")

	for i := 0; i < 10; i++ {
		system, err := rSystem.New(context.Background(), model.New(map[string]interface{}{
			"Name":   util.RandomString(10),
			"Tables": []string{},
		}))

		require.NoError(t, err)
		require.NotEmpty(t, system)
	}

	systems, err := rSystem.List(context.Background(), model.New(map[string]interface{}{
		"Limit":  int32(5),
		"Offset": int32(5),
	}))

	require.NoError(t, err)
	require.Len(t, systems, 5)

	for _, system := range systems {
		require.NotEmpty(t, system)
	}

	systems, err = rSystem.List(context.Background(), model.New(nil))

	require.NoError(t, err)
	require.Len(t, systems, 10)

	for _, system := range systems {
		require.NotEmpty(t, system)

		res, err := rSystem.Delete(context.Background(), model.New(map[string]interface{}{
			"ID": system.Get("ID").(uuid.UUID),
		}))

		require.NoError(t, err)
		require.NotEmpty(t, res)
	}
}

func TestUpdateSystem(t *testing.T) {
	rSystem := repo.Table("system")

	system, err := rSystem.New(context.Background(), model.New(map[string]interface{}{
		"Name":   util.RandomString(10),
		"Tables": []string{},
	}))

	require.NoError(t, err)
	require.NotEmpty(t, system)

	arg := model.New(map[string]interface{}{
		"ID":     system.Get("ID"),
		"Tables": []string{util.RandomString(10)},
		"Active": util.RandomBool(),
	})

	updated, err := rSystem.Update(context.Background(), arg)

	require.NoError(t, err)
	require.NotEmpty(t, system)

	require.Equal(t, system.Get("ID"), updated.Get("ID"))
	require.Equal(t, system.Get("Name"), updated.Get("Name"))
	require.Equal(t, arg.Get("Tables"), updated.Get("Tables"))
	require.Equal(t, arg.Get("Active"), updated.Get("Active"))

	system, err = rSystem.Delete(context.Background(), model.New(map[string]interface{}{
		"ID": system.Get("ID").(uuid.UUID),
	}))

	require.NoError(t, err)
	require.NotEmpty(t, system)
}
