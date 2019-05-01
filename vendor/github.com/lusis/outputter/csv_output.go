package outputter

import (
	"encoding/csv"
	"io"
	"os"
	"sync"
)

func init() {
	RegisterOutput("csv", newCSVFactoryOutput)
}

func newCSVFactoryOutput() Outputter {
	return NewCSVOutput()
}

// CSVOutput is an Outputter that draws data in tabular format
type CSVOutput struct {
	headers []string
	rows    [][]string
	writer  io.Writer
	sep     string
	sync.Mutex
}

// NewCSVOutput creates a New TablularOutput with os.Stdout as the destination
func NewCSVOutput() *CSVOutput {
	t := NewCSVOutputWithWriter(os.Stdout)
	return t
}

// NewCSVOutputWithWriter creates a new instance of TabularOutput with the provided io.Writer
func NewCSVOutputWithWriter(i io.Writer) *CSVOutput {
	t := &CSVOutput{}
	t.writer = i
	return t
}

// SetHeaders sets the table's headers
func (t *CSVOutput) SetHeaders(headers []string) {
	t.Lock()
	defer t.Unlock()
	t.headers = headers
	//
}

// AddRow adds a row to the table
func (t *CSVOutput) AddRow(row []string) error {
	t.Lock()
	defer t.Unlock()
	t.rows = append(t.rows, row)
	return nil
}

// SetPretty sets pretty output
func (t *CSVOutput) SetPretty() {
	//noop
}

// Draw displays the table to stdout
func (t *CSVOutput) Draw() {
	t.Lock()
	defer t.Unlock()
	w := csv.NewWriter(t.writer)

	_ = w.Write(t.headers)
	for _, row := range t.rows {
		/*
			var tmpRow []string
			for _, rec := range row {
				tmpRow = append(tmpRow, fmt.Sprintf("%s", rec))
			}
		*/
		_ = w.Write(row)
	}
	w.Flush()
}

// SetWriter sets the writer for output
func (t *CSVOutput) SetWriter(i io.Writer) error {
	t.Lock()
	defer t.Unlock()
	t.writer = i
	return nil
}

// ColorSupport specifies if the output supports colorized text or not
func (t *CSVOutput) ColorSupport() bool {
	// have to turn off color on tabular output for now =(
	// http://stackoverflow.com/questions/35398497/how-do-i-get-colors-to-work-with-golang-tabwriter
	return false
}
