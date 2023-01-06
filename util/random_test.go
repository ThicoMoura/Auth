package util_test

import (
	"testing"

	"github.com/ThicoMoura/Auth/util"
	"github.com/stretchr/testify/require"
)

func TestRandomInt(t *testing.T) {
	n := util.RandomInt(0, 10)

	require.GreaterOrEqual(t, n, int64(0))
	require.LessOrEqual(t, n, int64(10))
}

func TestRandomString(t *testing.T) {
	st := util.RandomString(10)

	require.NotEmpty(t, st)
	require.Len(t, st, 10)

	anotherSt := util.RandomString(10)

	require.NotEmpty(t, anotherSt)
	require.Len(t, anotherSt, 10)

	require.NotEqual(t, st, anotherSt)
}

func TestRandomBool(t *testing.T) {
	util.RandomBool()
}
