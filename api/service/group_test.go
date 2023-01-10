package service_test

import (
	"context"
	"testing"

	"github.com/ThicoMoura/Auth/api/model"
	"github.com/ThicoMoura/Auth/util"
	"github.com/google/uuid"
	"github.com/stretchr/testify/require"
)

func newGroup(t *testing.T) model.Model {
	group, err := services["group"].New(context.Background(), model.NewG{
		Name: util.RandomString(10),
	})

	require.NoError(t, err)
	require.NotEmpty(t, group)

	return group
}

func deleteGroup(t *testing.T, ID uuid.UUID) model.Model {
	group, err := services["group"].Delete(context.Background(), model.Id{
		ID: ID.String(),
	})

	require.NoError(t, err)
	require.NotEmpty(t, group)

	return group
}

func TestNewGroup(t *testing.T) {
	deleteGroup(t, newGroup(t).Get("ID").(uuid.UUID))
}

func TestGetGroup(t *testing.T) {
	group := newGroup(t)
	defer deleteGroup(t, group.Get("ID").(uuid.UUID))

	res, err := services["group"].Get(context.Background(), model.Id{
		ID: group.Get("ID").(uuid.UUID).String(),
	})

	require.NoError(t, err)
	require.NotEmpty(t, res)

	require.Equal(t, group, res)
}

func TestFindGroup(t *testing.T) {
	var IDs []uuid.UUID

	name := 0
	for i := 0; i < 10; i++ {
		group := newGroup(t)

		if group.Get("Name").(string)[0:1] == "a" {
			name++
		}

		IDs = append(IDs, group.Get("ID").(uuid.UUID))
	}

	defer func() {
		for _, id := range IDs {
			deleteGroup(t, id)
		}
	}()

	res, err := services["group"].Find(context.Background(), model.FindG{
		Name: "a",
	})

	require.NoError(t, err)
	require.Len(t, res, name)

	for _, group := range res {
		require.NotEmpty(t, group)
	}

	id := int32(1)
	size := int32(1)

	res, err = services["group"].Find(context.Background(), model.FindG{
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

	for _, group := range res {
		require.NotEmpty(t, group)
	}
}

func TestListGroup(t *testing.T) {
	for i := 0; i < 10; i++ {
		newGroup(t)
	}

	id := int32(2)
	size := int32(5)

	list, err := services["group"].List(context.Background(), model.List{
		PageID:   &id,
		PageSize: &size,
	})

	require.NoError(t, err)
	require.Len(t, list, 5)

	for _, group := range list {
		require.NotEmpty(t, group)
	}

	list, err = services["group"].List(context.Background(), model.List{})

	require.NoError(t, err)
	require.Len(t, list, 10)

	for _, group := range list {
		require.NotEmpty(t, group)
		deleteGroup(t, group.Get("ID").(uuid.UUID))
	}
}

func TestUpdateGroup(t *testing.T) {
	group := newGroup(t)
	defer deleteGroup(t, group.Get("ID").(uuid.UUID))

	name := util.RandomString(10)

	res, err := services["group"].Update(context.Background(), model.UpdateG{
		ID:   group.Get("ID").(uuid.UUID).String(),
		Name: &name,
	})

	require.NoError(t, err)
	require.NotEmpty(t, res)

	require.Equal(t, group.Get("ID"), res.Get("ID"))
	require.Equal(t, name, res.Get("Name"))
	require.Equal(t, group.Get("Active"), res.Get("Active"))

	active := util.RandomBool()

	res, err = services["group"].Update(context.Background(), model.UpdateG{
		ID:     group.Get("ID").(uuid.UUID).String(),
		Active: &active,
	})

	require.NoError(t, err)
	require.NotEmpty(t, res)

	require.Equal(t, group.Get("ID"), res.Get("ID"))
	require.Equal(t, name, res.Get("Name"))
	require.Equal(t, active, res.Get("Active"))
}
