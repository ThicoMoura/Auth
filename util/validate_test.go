package util_test

import (
	"testing"

	"github.com/ThicoMoura/Auth/util"
	"github.com/stretchr/testify/require"
)

type test struct {
	Field string `validate:"required"`
}

func TestRequired(t *testing.T) {
	test := test{}

	require.NotEmpty(t, util.NewValidate().Validator(test))

	for key, value := range util.NewValidate().Validator(test) {
		require.Equal(t, "Field", key)
		require.Equal(t, "required", value)
	}

	test.Field = "Test"

	require.Empty(t, util.NewValidate().Validator(test))
}
