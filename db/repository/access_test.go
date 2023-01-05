package repository_test

import (
	"context"
	"testing"

	"github.com/ThicoMoura/Auth/db/model"
	"github.com/ThicoMoura/Auth/util"
	"github.com/google/uuid"
	"github.com/stretchr/testify/require"
)

func NewAccess(t *testing.T, ID uuid.UUID) model.Value {
	access, err := repo.Table("access").New(context.Background(), model.New(map[string]interface{}{
		"System": ID,
		"Table":  util.RandomString(10),
		"Type":   []string{util.RandomString(10), util.RandomString(10)},
	}))

	require.NoError(t, err)
	require.NotEmpty(t, access)

	return access
}

func DeleteAccess(t *testing.T, ID uuid.UUID) model.Value {
	access, err := repo.Table("access").Delete(context.Background(), model.New(map[string]interface{}{
		"ID": ID,
	}))

	require.NoError(t, err)
	require.NotEmpty(t, access)

	return access
}

func TestNewAccess(t *testing.T) {
	access := NewAccess(t, NewSystem(t).Get("ID").(uuid.UUID))
	DeleteAccess(t, access.Get("ID").(uuid.UUID))
	DeleteSystem(t, access.Get("System").(uuid.UUID))
}

func TestGetAccess(t *testing.T) {
	access := NewAccess(t, NewSystem(t).Get("ID").(uuid.UUID))

	res, err := repo.Table("access").Get(context.Background(), model.New(map[string]interface{}{
		"ID": access.Get("ID").(uuid.UUID),
	}))

	require.NoError(t, err)
	require.NotEmpty(t, res)

	require.Equal(t, access, res)

	DeleteAccess(t, access.Get("ID").(uuid.UUID))
	DeleteSystem(t, access.Get("System").(uuid.UUID))
}

func TestFindAccess(t *testing.T) {
	var accesss []model.Value

	system := NewSystem(t)
	table := 0

	for i := 0; i < 10; i++ {
		access := NewAccess(t, system.Get("ID").(uuid.UUID))

		if access.Get("Table").(string)[0:1] == "a" {
			table++
		}

		accesss = append(accesss, access)
	}

	res, err := repo.Table("access").Find(context.Background(), model.New(map[string]interface{}{
		"Table": "a%",
	}))

	require.NoError(t, err)
	require.Len(t, res, table)

	for _, access := range res {
		require.NotEmpty(t, access)
	}

	for _, access := range accesss {
		DeleteAccess(t, access.Get("ID").(uuid.UUID))
	}

	DeleteSystem(t, system.Get("ID").(uuid.UUID))
}

func TestListAccess(t *testing.T) {
	system := NewSystem(t)

	for i := 0; i < 10; i++ {
		NewAccess(t, system.Get("ID").(uuid.UUID))
	}

	accesss, err := repo.Table("access").List(context.Background(), model.New(map[string]interface{}{
		"Limit":  int32(5),
		"Offset": int32(5),
	}))

	require.NoError(t, err)
	require.Len(t, accesss, 5)

	for _, access := range accesss {
		require.NotEmpty(t, access)
	}

	accesss, err = repo.Table("access").List(context.Background(), model.New(nil))

	require.NoError(t, err)
	require.Len(t, accesss, 10)

	for _, access := range accesss {
		require.NotEmpty(t, access)

		DeleteAccess(t, access.Get("ID").(uuid.UUID))
	}

	DeleteSystem(t, system.Get("ID").(uuid.UUID))
}

func TestUpdateAccess(t *testing.T) {
	access := NewAccess(t, NewSystem(t).Get("ID").(uuid.UUID))

	arg := model.New(map[string]interface{}{
		"ID":   access.Get("ID"),
		"Type": []string{util.RandomString(10)},
	})

	updated, err := repo.Table("access").Update(context.Background(), arg)

	require.NoError(t, err)
	require.NotEmpty(t, access)

	require.Equal(t, access.Get("ID"), updated.Get("ID"))
	require.Equal(t, arg.Get("Type"), updated.Get("Type"))

	DeleteAccess(t, access.Get("ID").(uuid.UUID))
	DeleteSystem(t, access.Get("System").(uuid.UUID))
}
