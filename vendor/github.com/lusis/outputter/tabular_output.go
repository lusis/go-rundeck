package outputter

import (
	"fmt"
	"io"
	"os"
	"strings"
	"sync"
	"text/tabwriter"
)

func init() {
	RegisterOutput("tabular", newTabularFactoryOutput)
}

func newTabularFactoryOutput() Outputter {
	return NewTabularOutput()
}

// TabularOutput is an Outputter that draws data in tabular format
type TabularOutput struct {
	table   *tabwriter.Writer
	headers []string
	rows    [][]string
	writer  io.Writer
	sync.Mutex
}

// NewTabularOutput creates a New TablularOutput with os.Stdout as the destination
func NewTabularOutput() *TabularOutput {
	t := NewTabularOutputWithWriter(os.Stdout)
	return t
}

// NewTabularOutputWithWriter creates a new instance of TabularOutput with the provided io.Writer
func NewTabularOutputWithWriter(i io.Writer) *TabularOutput {
	t := &TabularOutput{}
	t.writer = i
	return t
}

// SetHeaders sets the table's headers
func (t *TabularOutput) SetHeaders(headers []string) {
	t.Lock()
	defer t.Unlock()
	t.headers = headers
	//
}

// AddRow adds a row to the table
func (t *TabularOutput) AddRow(row []string) error {
	t.Lock()
	defer t.Unlock()
	t.rows = append(t.rows, row)
	return nil
}

// SetPretty sets pretty output
func (t *TabularOutput) SetPretty() {
	//noop
}

// Draw displays the table to stdout
func (t *TabularOutput) Draw() {
	t.Lock()
	defer t.Unlock()
	w := tabwriter.NewWriter(t.writer, 0, 2, 2, ' ', 0)
	fmt.Fprintf(w, strings.Join(t.headers, "\t")+"\n")
	for _, row := range t.rows {
		fmt.Fprintf(w, strings.Join(row, "\t")+"\n")
	}
	_ = w.Flush()
}

// SetWriter sets the writer for output
func (t *TabularOutput) SetWriter(i io.Writer) error {
	t.Lock()
	defer t.Unlock()
	t.writer = i
	return nil
}

// ColorSupport specifies if the output supports colorized text or not
func (t *TabularOutput) ColorSupport() bool {
	// have to turn off color on tabular output for now =(
	// http://stackoverflow.com/questions/35398497/how-do-i-get-colors-to-work-with-golang-tabwriter
	return false
}
