package main

import (
	"bufio"
	"bytes"
	"testing"

	"github.com/lusis/outputter"

	"github.com/stretchr/testify/assert"
)

func TestNewCustomOutput(t *testing.T) {
	o := NewCustomOutput()
	assert.IsType(t, NewCustomOutput(), o)
}

func TestRegisterCustomOutput(t *testing.T) {
	outputter.RegisterOutput("testoutput", NewCustomFactoryOutput)
	allOutputs := outputter.GetOutputters()
	assert.Contains(t, allOutputs, "testoutput")
}

func TestCustomOutputSingleRow(t *testing.T) {
	var buf bytes.Buffer
	output := bufio.NewWriter(&buf)
	o := NewCustomOutputWithWriter(output)
	o.SetHeaders([]string{"key1", "key2"})
	err := o.AddRow([]string{"value1", "value2"})
	assert.NoError(t, err)
	o.Draw()
	_ = output.Flush()
	expected := "here's my data!\nkey1:value1\nkey2:value2\n"
	assert.Equal(t, expected, buf.String())
}
func TestCustomOutputMultipleRows(t *testing.T) {
	var buf bytes.Buffer
	output := bufio.NewWriter(&buf)
	o := NewCustomOutputWithWriter(output)
	o.SetHeaders([]string{"key1", "key2"})
	err1 := o.AddRow([]string{"value1", "value2"})
	assert.NoError(t, err1)
	err2 := o.AddRow([]string{"value3", "value4"})
	assert.NoError(t, err2)
	o.Draw()
	_ = output.Flush()
	expected := "here's my data!\nkey1:value1,value3\nkey2:value2,value4\n"
	assert.Equal(t, expected, buf.String())
}

func TestCustomOutputMissingKeys(t *testing.T) {
	outputter.RegisterOutput("testoutput", NewCustomFactoryOutput)
	o, oErr := outputter.NewOutputter("testoutput")
	assert.NoError(t, oErr)
	err := o.AddRow([]string{"key1", "key2"})
	assert.Equal(t, outputter.ErrorOutputAddRowNoHeaders, err)
}

func TestCustomOutputToFewValues(t *testing.T) {
	outputter.RegisterOutput("testoutput", NewCustomFactoryOutput)
	o, oErr := outputter.NewOutputter("testoutput")
	assert.NoError(t, oErr)
	o.SetHeaders([]string{"key1"})
	err := o.AddRow([]string{"value1", "value2"})
	assert.Equal(t, outputter.ErrorOutputAddRowTooFewHeaders, err)
}

func TestCustomOutputFewerValues(t *testing.T) {
	var buf bytes.Buffer
	output := bufio.NewWriter(&buf)
	o := NewCustomOutputWithWriter(output)
	o.SetHeaders([]string{"key1", "key2"})
	r1Err := o.AddRow([]string{"value1", "value2"})
	r2Err := o.AddRow([]string{"value3"})
	assert.NoError(t, r1Err)
	assert.NoError(t, r2Err)
	o.Draw()
	_ = output.Flush()
	assert.Equal(t, "here's my data!\nkey1:value1,value3\nkey2:value2,\n", buf.String())
}
