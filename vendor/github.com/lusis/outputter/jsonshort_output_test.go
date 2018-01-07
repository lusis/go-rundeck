package outputter

import (
	"testing"

	"bufio"
	"bytes"

	"github.com/stretchr/testify/assert"
)

func TestNewJSONShortOutput(t *testing.T) {
	j := NewJSONShortOutput()
	assert.IsType(t, NewJSONShortOutput(), j)
}

func TestJSONShortOutputSingleRow(t *testing.T) {
	var buf bytes.Buffer
	output := bufio.NewWriter(&buf)
	json := NewJSONShortOutputWithWriter(output)
	json.SetHeaders([]string{"key1", "key2"})
	err := json.AddRow([]string{"value1", "value2"})
	assert.NoError(t, err)
	json.Draw()
	_ = output.Flush()
	assert.Equal(t, `{"key1":["value1"],"key2":["value2"]}`, buf.String())
}

func TestJSONShortOutputSetWriter(t *testing.T) {
	var buf bytes.Buffer
	output := bufio.NewWriter(&buf)
	json := NewJSONShortOutput()
	jsonErr := json.SetWriter(output)
	assert.NoError(t, jsonErr)
	json.SetHeaders([]string{"key1", "key2"})
	err := json.AddRow([]string{"value1", "value2"})
	assert.NoError(t, err)
	json.Draw()
	_ = output.Flush()
	assert.Equal(t, `{"key1":["value1"],"key2":["value2"]}`, buf.String())
}

func TestJSONShortOutputMultipleRows(t *testing.T) {
	var buf bytes.Buffer
	json := NewJSONShortOutputWithWriter(&buf)
	json.SetHeaders([]string{"key1", "key2"})
	r1Err := json.AddRow([]string{"value1", "value2"})
	r2Err := json.AddRow([]string{"value3", "value4"})
	assert.NoError(t, r1Err)
	assert.NoError(t, r2Err)
	json.Draw()
	assert.Equal(t, `{"key1":["value1","value3"],"key2":["value2","value4"]}`, buf.String())
}

func TestJSONShortOutputPretty(t *testing.T) {
	var buf bytes.Buffer
	json := NewJSONShortOutputWithWriter(&buf)
	json.SetPretty()
	json.SetHeaders([]string{"key1", "key2"})
	r1Err := json.AddRow([]string{"value1", "value2"})
	r2Err := json.AddRow([]string{"value3", "value4"})
	assert.NoError(t, r1Err)
	assert.NoError(t, r2Err)
	json.Draw()
	assert.Equal(t, "{\n\t\"key1\": [\n\t\t\"value1\",\n\t\t\"value3\"\n\t],\n\t\"key2\": [\n\t\t\"value2\",\n\t\t\"value4\"\n\t]\n}\n", buf.String())
}

func TestJSONShortOutputMissingKeys(t *testing.T) {
	j := NewJSONShortOutput()
	err := j.AddRow([]string{"key1", "key2"})
	assert.Equal(t, ErrorOutputAddRowNoHeaders, err)
}

func TestJSONShortOutputToFewValues(t *testing.T) {
	j := NewJSONShortOutput()
	j.SetHeaders([]string{"key1"})
	err := j.AddRow([]string{"value1", "value2"})
	assert.Equal(t, ErrorOutputAddRowTooFewHeaders, err)
}

func TestJSONShortOutputFewerValues(t *testing.T) {
	var buf bytes.Buffer
	output := bufio.NewWriter(&buf)
	json := NewJSONShortOutputWithWriter(output)
	json.SetHeaders([]string{"key1", "key2"})
	r1Err := json.AddRow([]string{"value1", "value2"})
	r2Err := json.AddRow([]string{"value3"})
	assert.NoError(t, r1Err)
	assert.NoError(t, r2Err)
	json.Draw()
	_ = output.Flush()
	assert.Equal(t, `{"key1":["value1","value3"],"key2":["value2",""]}`, buf.String())
}
