package outputter

import (
	"bytes"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewTabularOutput(t *testing.T) {
	table := NewTabularOutput()
	assert.IsType(t, new(TabularOutput), table)
}

func TestTabularOutputNew(t *testing.T) {
	var buf bytes.Buffer
	table := NewTabularOutputWithWriter(&buf)
	table.SetHeaders([]string{"header1", "header2"})
	r1Err := table.AddRow([]string{"value1aaaaaaaaaaaaa", "value2bbbbbbbbbbbbbbbbbb"})
	r2Err := table.AddRow([]string{"value3", "value4"})
	assert.NoError(t, r1Err)
	assert.NoError(t, r2Err)
	table.Draw()
	expectedOutput := "header1              header2\nvalue1aaaaaaaaaaaaa  value2bbbbbbbbbbbbbbbbbb\nvalue3               value4\n"
	assert.Equal(t, expectedOutput, buf.String())
}

func TestTabularOutputSetWriter(t *testing.T) {
	var buf bytes.Buffer
	table := NewTabularOutput()
	tErr := table.SetWriter(&buf)
	assert.NoError(t, tErr)
	table.SetHeaders([]string{"header1", "header2"})
	r1Err := table.AddRow([]string{"value1aaaaaaaaaaaaa", "value2bbbbbbbbbbbbbbbbbb"})
	r2Err := table.AddRow([]string{"value3", "value4"})
	assert.NoError(t, r1Err)
	assert.NoError(t, r2Err)
	table.Draw()
	expectedOutput := "header1              header2\nvalue1aaaaaaaaaaaaa  value2bbbbbbbbbbbbbbbbbb\nvalue3               value4\n"
	assert.Equal(t, expectedOutput, buf.String())
}
