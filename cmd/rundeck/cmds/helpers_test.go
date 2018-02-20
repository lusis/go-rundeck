package cmds

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestParseGood(t *testing.T) {
	data := []string{"foo=bar", "qux=baz"}
	res, err := ParseSliceKeyValue(data)
	require.NoError(t, err)
	require.Len(t, res, 2)
}

func TestParseMissingKey(t *testing.T) {
	data := []string{"=bar", "qux=baz"}
	res, err := ParseSliceKeyValue(data)
	require.Error(t, err)
	require.Len(t, res, 0)
}

func TestParseMissingValue(t *testing.T) {
	data := []string{"foo=", "qux=baz"}
	res, err := ParseSliceKeyValue(data)
	require.Error(t, err)
	require.Len(t, res, 0)
}

func TestParseNoEqual(t *testing.T) {
	data := []string{"foo", "qux=baz"}
	res, err := ParseSliceKeyValue(data)
	require.Error(t, err)
	require.Len(t, res, 0)
}
