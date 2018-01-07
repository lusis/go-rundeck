package outputter

import (
	"io"
	"os"
	"sync"

	"github.com/olekukonko/tablewriter"
)

func init() {
	RegisterOutput("table", newTableFactoryOutput)
}

// TableOutput is an Outputter that draw data as a table
type TableOutput struct {
	table   *tablewriter.Table
	writer  io.Writer
	headers []string
	rows    [][]string
	sync.Mutex
}

func newTableFactoryOutput() Outputter {
	x := NewTableOutput()
	return x
}

// NewTableOutput creates a New TableOutput with os.Stdout as the destination
func NewTableOutput() *TableOutput {
	return NewTableOutputWithWriter(os.Stdout)
}

// NewTableOutputWithWriter creates a new instance of TableOutput with the provided io.Writer
func NewTableOutputWithWriter(i io.Writer) *TableOutput {
	t := &TableOutput{}
	t.writer = i
	return t
}

// SetHeaders sets the table's headers
func (t *TableOutput) SetHeaders(headers []string) {
	t.Lock()
	defer t.Unlock()
	t.headers = headers
	//t.table.SetHeader(headers)
}

// AddRow adds a row to the table
func (t *TableOutput) AddRow(row []string) error {
	t.Lock()
	defer t.Unlock()
	if len(t.headers) == 0 {
		return ErrorOutputAddRowNoHeaders
	}
	if len(t.headers) < len(row) {
		return ErrorOutputAddRowTooFewHeaders
	}
	t.rows = append(t.rows, row)
	return nil
}

// Draw displays the table to stdout
func (t *TableOutput) Draw() {
	t.Lock()
	defer t.Unlock()
	tw := tablewriter.NewWriter(t.writer)
	tw.SetAutoWrapText(false)
	tw.SetHeaderAlignment(tablewriter.ALIGN_LEFT)
	tw.SetAlignment(tablewriter.ALIGN_LEFT)
	tw.SetHeader(t.headers)
	tw.AppendBulk(t.rows)
	tw.Render()
}

// SetPretty returns a prettified version
func (t *TableOutput) SetPretty() {
	//noop for table
}

// SetWriter sets the io.Writer.
func (t *TableOutput) SetWriter(i io.Writer) error {
	t.Lock()
	defer t.Unlock()
	t.writer = i
	return nil
}

// ColorSupport specifies if the output supports colorized text or not
func (t *TableOutput) ColorSupport() bool {
	return true
}
