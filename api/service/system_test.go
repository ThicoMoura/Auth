package service_test

import (
	"context"
	"testing"

	"github.com/ThicoMoura/Auth/api/model"
	"github.com/ThicoMoura/Auth/util"
	"github.com/google/uuid"
	"github.com/stretchr/testify/require"
)

func newSystem(t *testing.T) model.Model {
	system, err := services["system"].New(context.Background(), model.NewSy{
		Name: util.RandomString(10),
	})

	require.NoError(t, err)
	require.NotEmpty(t, system)

	return system
}

func deleteSystem(t *testing.T, ID uuid.UUID) model.Model {
	system, err := services["system"].Delete(context.Background(), model.Id{
		ID: ID.String(),
	})

	require.NoError(t, err)
	require.NotEmpty(t, system)

	return system
}

func TestNewSystem(t *testing.T) {
	defer deleteSystem(t, newSystem(t).Get("ID").(uuid.UUID))
}

func TestGetSystem(t *testing.T) {
	system := newSystem(t)

	defer deleteSystem(t, system.Get("ID").(uuid.UUID))

	res, err := services["system"].Get(context.Background(), model.Id{
		ID: system.Get("ID").(uuid.UUID).String(),
	})

	require.NoError(t, err)
	require.NotEmpty(t, res)

	require.Equal(t, system, res)
}

func TestFindSystem(t *testing.T) {
	var IDs []uuid.UUID

	name := 0
	for i := 0; i < 10; i++ {
		system := newSystem(t)

		if system.Get("Name").(string)[0:1] == "a" {
			name++
		}

		IDs = append(IDs, system.Get("ID").(uuid.UUID))
	}

	res, err := services["system"].Find(context.Background(), model.FindSy{
		Name: "a",
	})

	require.NoError(t, err)
	require.Len(t, res, name)

	for _, system := range res {
		require.NotEmpty(t, system)
	}

	id := int32(1)
	size := int32(1)

	res, err = services["system"].Find(context.Background(), model.FindSy{
		Name:     "a",
		PageID:   &id,
		PageSize: &size,
	})

	result := name
	if result > 1 {
		result = 1
	}

	require.NoError(t, err)
	require.Len(t, res, result)

	for _, system := range res {
		require.NotEmpty(t, system)
	}

	for _, id := range IDs {
		deleteSystem(t, id)
	}
}

func TestListSystem(t *testing.T) {
	for i := 0; i < 10; i++ {
		newSystem(t)
	}

	id := int32(2)
	size := int32(5)

	list, err := services["system"].List(context.Background(), model.List{
		PageID:   &id,
		PageSize: &size,
	})

	require.NoError(t, err)
	require.Len(t, list, 5)

	for _, system := range list {
		require.NotEmpty(t, system)
	}

	list, err = services["system"].List(context.Background(), model.List{})

	require.NoError(t, err)
	require.Len(t, list, 10)

	for _, system := range list {
		require.NotEmpty(t, system)
		deleteSystem(t, system.Get("ID").(uuid.UUID))
	}
}

func TestUpdateSystem(t *testing.T) {
	system := newSystem(t)

	name := util.RandomString(10)

	res, err := services["system"].Update(context.Background(), model.UpdateSy{
		ID:   system.Get("ID").(uuid.UUID).String(),
		Name: &name,
	})

	require.NoError(t, err)
	require.NotEmpty(t, res)

	require.Equal(t, system.Get("ID"), res.Get("ID"))
	require.Equal(t, name, res.Get("Name"))
	require.Equal(t, system.Get("Active"), res.Get("Active"))

	active := util.RandomBool()

	res, err = services["system"].Update(context.Background(), model.UpdateSy{
		ID:     system.Get("ID").(uuid.UUID).String(),
		Active: &active,
	})

	require.NoError(t, err)
	require.NotEmpty(t, res)

	require.Equal(t, system.Get("ID"), res.Get("ID"))
	require.Equal(t, name, res.Get("Name"))
	require.Equal(t, active, res.Get("Active"))

	deleteSystem(t, system.Get("ID").(uuid.UUID))
}
