package outputter

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewOutputterJSON(t *testing.T) {
	f, err := NewOutputter("json")
	assert.NoError(t, err)
	assert.IsType(t, &JSONOutput{}, f)
}

func TestGetDefaultOutputter(t *testing.T) {
	f := GetDefaultOutputter()
	assert.IsType(t, &TabularOutput{}, f)
}

func TestKnownOutputs(t *testing.T) {
	f := GetOutputters()
	assert.NotEmpty(t, f)
	assert.Len(t, f, 5)
}

func TestUnknownOutputter(t *testing.T) {
	_, err := NewOutputter("test")
	assert.Equal(t, err, ErrorUnknownOutputter)
}

func TestNewKnownOutputters(t *testing.T) {
	for _, x := range GetOutputters() {
		_, err := NewOutputter(x)
		assert.NoError(t, err)
	}
}
