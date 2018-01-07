package outputter

import (
	"bytes"
	"testing"

	"github.com/fatih/color"
	"github.com/stretchr/testify/assert"
)

func TestNewTableOutput(t *testing.T) {
	table := NewTableOutput()
	assert.IsType(t, new(TableOutput), table)
}

func TestTableOutputNew(t *testing.T) {
	var buf bytes.Buffer
	table := NewTableOutputWithWriter(&buf)
	table.SetHeaders([]string{"header1", "header2"})
	err := table.AddRow([]string{"value1", "value2"})
	assert.NoError(t, err)
	table.Draw()
	expectedOutput := `+---------+---------+
| HEADER1 | HEADER2 |
+---------+---------+
| value1  | value2  |
+---------+---------+
`
	assert.Equal(t, expectedOutput, buf.String())
}

func TestTableTooManyValues(t *testing.T) {
	table := NewTableOutput()
	table.SetHeaders([]string{"header1", "header2"})
	err := table.AddRow([]string{"value1", "value2", "value3"})
	assert.Equal(t, ErrorOutputAddRowTooFewHeaders, err)
}

func TestTableNoHeaders(t *testing.T) {
	table := NewTableOutput()
	err := table.AddRow([]string{"value1", "value2", "value3"})
	assert.Equal(t, ErrorOutputAddRowNoHeaders, err)
}
func TestTableOutputFewerValues(t *testing.T) {
	var buf bytes.Buffer
	table := NewTableOutputWithWriter(&buf)
	table.SetHeaders([]string{"header1", "header2"})
	err := table.AddRow([]string{"value1"})
	assert.NoError(t, err)
	table.Draw()
	expectedOutput := `+---------+---------+
| HEADER1 | HEADER2 |
+---------+---------+
| value1  |
+---------+---------+
`
	assert.Equal(t, expectedOutput, buf.String())
}

func TestTableOutputChangeWriter(t *testing.T) {
	var buf1 bytes.Buffer
	var buf2 bytes.Buffer
	table := NewTableOutputWithWriter(&buf1)
	table.SetHeaders([]string{"header1", "header2"})
	err := table.AddRow([]string{"value1", "value2"})
	assert.NoError(t, err)
	table.Draw()
	expectedOutput := `+---------+---------+
| HEADER1 | HEADER2 |
+---------+---------+
| value1  | value2  |
+---------+---------+
`
	assert.Equal(t, expectedOutput, buf1.String())
	wErr := table.SetWriter(&buf2)
	assert.NoError(t, wErr)
	table.Draw()
	assert.Equal(t, expectedOutput, buf2.String())
}

func TestTableOutputSetPrettyNoop(t *testing.T) {
	var buf1 bytes.Buffer
	var buf2 bytes.Buffer
	table := NewTableOutputWithWriter(&buf1)
	table.SetHeaders([]string{"header1", "header2"})
	err := table.AddRow([]string{"value1", "value2"})
	assert.NoError(t, err)
	table.Draw()
	expectedOutput := `+---------+---------+
| HEADER1 | HEADER2 |
+---------+---------+
| value1  | value2  |
+---------+---------+
`
	assert.Equal(t, expectedOutput, buf1.String())
	wErr := table.SetWriter(&buf2)
	assert.NoError(t, wErr)
	table.SetPretty()
	table.Draw()
	assert.Equal(t, expectedOutput, buf2.String())
}

func TestTableOutputColorized(t *testing.T) {
	var buf1 bytes.Buffer
	table := NewTableOutputWithWriter(&buf1)
	color.NoColor = !table.ColorSupport()
	table.SetHeaders([]string{color.RedString("header1"), color.RedString("header2")})
	err := table.AddRow([]string{color.BlueString("value1"), color.BlueString("value2")})
	assert.NoError(t, err)
	table.Draw()
	expectedOutput := "+---------+---------+\n| \x1b[31MHEADER1\x1b[0M | \x1b[31MHEADER2\x1b[0M |\n+---------+---------+\n| \x1b[34mvalue1\x1b[0m  | \x1b[34mvalue2\x1b[0m  |\n+---------+---------+\n"

	assert.Equal(t, expectedOutput, buf1.String())
}
