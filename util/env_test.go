package util_test

import (
	"testing"

	"github.com/ThicoMoura/Auth/util"
	"github.com/stretchr/testify/require"
)

func TestNewEnv(t *testing.T) {
	env, err := util.NewEnv("../")

	require.NoError(t, err)
	require.NotEmpty(t, env)
}
