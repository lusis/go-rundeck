package outputter

import (
	"bufio"
	"bytes"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewCSVOutput(t *testing.T) {
	j := NewCSVOutput()
	assert.IsType(t, NewCSVOutput(), j)
}

func TestCSVOutputSingleRow(t *testing.T) {
	var buf bytes.Buffer
	output := bufio.NewWriter(&buf)
	csv := NewCSVOutputWithWriter(output)
	csv.SetHeaders([]string{"key1", "key2"})
	err := csv.AddRow([]string{"value1", "value2"})
	assert.NoError(t, err)
	csv.Draw()
	_ = output.Flush()
	assert.Equal(t, "key1,key2\nvalue1,value2\n", buf.String())
}

func TestCSVOutputMultipleRows(t *testing.T) {
	var buf bytes.Buffer
	csv := NewCSVOutputWithWriter(&buf)
	csv.SetHeaders([]string{"key1", "key2"})
	r1Err := csv.AddRow([]string{"value1", "value2"})
	r2Err := csv.AddRow([]string{"value3", "value4"})
	assert.NoError(t, r1Err)
	assert.NoError(t, r2Err)
	csv.Draw()
	assert.Equal(t, "key1,key2\nvalue1,value2\nvalue3,value4\n", buf.String())
}

func TestCSVOutputSetWriter(t *testing.T) {
	var buf bytes.Buffer
	csv := NewCSVOutput()
	setErr := csv.SetWriter(&buf)
	assert.NoError(t, setErr)
	csv.SetHeaders([]string{"key1", "key2"})
	r1Err := csv.AddRow([]string{"value1", "value2"})
	r2Err := csv.AddRow([]string{"value3", "value4"})
	assert.NoError(t, r1Err)
	assert.NoError(t, r2Err)
	csv.Draw()
	assert.Equal(t, "key1,key2\nvalue1,value2\nvalue3,value4\n", buf.String())
}
