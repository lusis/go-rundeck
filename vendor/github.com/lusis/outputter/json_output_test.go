package outputter

import (
	"testing"

	"bufio"
	"bytes"

	"github.com/stretchr/testify/assert"
)

func TestNewJSONOutput(t *testing.T) {
	j := NewJSONOutput()
	assert.IsType(t, NewJSONOutput(), j)
}

func TestJSONOutputSingleRow(t *testing.T) {
	var buf bytes.Buffer
	output := bufio.NewWriter(&buf)
	json := NewJSONOutputWithWriter(output)
	json.SetHeaders([]string{"key1", "key2"})
	err := json.AddRow([]string{"value1", "value2"})
	assert.NoError(t, err)
	json.Draw()
	_ = output.Flush()
	assert.Equal(t, `[{"key1":"value1","key2":"value2"}]`, buf.String())
}

func TestJSONOutputSetWriter(t *testing.T) {
	var buf bytes.Buffer
	output := bufio.NewWriter(&buf)
	json := NewJSONOutput()
	setErr := json.SetWriter(&buf)
	assert.NoError(t, setErr)
	json.SetHeaders([]string{"key1", "key2"})
	err := json.AddRow([]string{"value1", "value2"})
	assert.NoError(t, err)
	json.Draw()
	_ = output.Flush()
	assert.Equal(t, `[{"key1":"value1","key2":"value2"}]`, buf.String())
}

func TestJSONOutputMultipleRows(t *testing.T) {
	var buf bytes.Buffer
	json := NewJSONOutputWithWriter(&buf)
	json.SetHeaders([]string{"key1", "key2"})
	r1Err := json.AddRow([]string{"value1", "value2"})
	r2Err := json.AddRow([]string{"value3", "value4"})
	assert.NoError(t, r1Err)
	assert.NoError(t, r2Err)
	json.Draw()
	assert.Equal(t, `[{"key1":"value1","key2":"value2"},{"key1":"value3","key2":"value4"}]`, buf.String())
}

func TestJSONOutputPretty(t *testing.T) {
	var buf bytes.Buffer
	json := NewJSONOutputWithWriter(&buf)
	json.SetPretty()
	json.SetHeaders([]string{"key1", "key2"})
	r1Err := json.AddRow([]string{"value1", "value2"})
	r2Err := json.AddRow([]string{"value3", "value4"})
	assert.NoError(t, r1Err)
	assert.NoError(t, r2Err)
	json.Draw()
	assert.Equal(t, "[\n\t{\n\t\t\"key1\": \"value1\",\n\t\t\"key2\": \"value2\"\n\t},\n\t{\n\t\t\"key1\": \"value3\",\n\t\t\"key2\": \"value4\"\n\t}\n]\n", buf.String())
}

func TestJSONOutputMissingKeys(t *testing.T) {
	j := NewJSONOutput()
	err := j.AddRow([]string{"key1", "key2"})
	assert.Equal(t, ErrorOutputAddRowNoHeaders, err)
}

func TestJSONOutputToFewValues(t *testing.T) {
	j := NewJSONOutput()
	j.SetHeaders([]string{"key1"})
	err := j.AddRow([]string{"value1", "value2"})
	assert.Equal(t, ErrorOutputAddRowTooFewHeaders, err)
}

func TestJSONOutputFewerValues(t *testing.T) {
	var buf bytes.Buffer
	output := bufio.NewWriter(&buf)
	json := NewJSONOutputWithWriter(output)
	json.SetHeaders([]string{"key1", "key2"})
	r1Err := json.AddRow([]string{"value1", "value2"})
	r2Err := json.AddRow([]string{"value3"})
	assert.NoError(t, r1Err)
	assert.NoError(t, r2Err)
	json.Draw()
	_ = output.Flush()
	assert.Equal(t, `[{"key1":"value1","key2":"value2"},{"key1":"value3","key2":""}]`, buf.String())
}
