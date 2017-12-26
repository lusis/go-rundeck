package outputter

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"os"
	"sync"
)

func init() {
	RegisterOutput("jsonshort", newJSONShortFactoryOutput)
}

// JSONShortOutput is an Outputter that draws data as json like so:
// {"header1":["value1","value3], "header2":["value2","value4"]}
type JSONShortOutput struct {
	pretty bool
	keys   []string
	writer io.Writer
	rows   [][]string
	sync.RWMutex
}

func newJSONShortFactoryOutput() Outputter {
	return NewJSONShortOutput()
}

// NewJSONShortOutput creates a new JSONOutput with os.Stdout as the destination
func NewJSONShortOutput() *JSONShortOutput {
	return NewJSONShortOutputWithWriter(os.Stdout)
}

// NewJSONShortOutputWithWriter returns a new JSONOutput with the specified io.Writer
func NewJSONShortOutputWithWriter(i io.Writer) *JSONShortOutput {
	t := &JSONShortOutput{}
	t.writer = i
	return t
}

// SetHeaders sets the JSON keys to be used
func (t *JSONShortOutput) SetHeaders(headers []string) {
	t.keys = headers
}

// AddRow adds a new set of data to the JSON array
func (t *JSONShortOutput) AddRow(row []string) error {
	if len(t.keys) == 0 {
		return ErrorOutputAddRowNoHeaders
	}
	if len(t.keys) < len(row) {
		return ErrorOutputAddRowTooFewHeaders
	}
	t.Lock()
	defer t.Unlock()
	if len(row) < len(t.keys) {
		difference := len(t.keys) - len(row)
		// we have to account for this and fill in empty values
		missingVals := make([]string, difference)
		row = append(row, missingVals...)
	}
	t.rows = append(t.rows, row)
	return nil
}

// SetPretty sets json output to pretty format
func (t *JSONShortOutput) SetPretty() {
	t.Lock()
	defer t.Unlock()
	t.pretty = true
}

// SetWriter sets the output io.Writer
func (t *JSONShortOutput) SetWriter(i io.Writer) error {
	t.Lock()
	defer t.Unlock()
	t.writer = i
	return nil
}

// Draw renders the table to the configured io.Writer
func (t *JSONShortOutput) Draw() {
	t.RLock()
	defer t.RUnlock()

	m := make(map[string][]string)
	for keyIdx, keyName := range t.keys {
		for _, row := range t.rows {
			m[keyName] = append(m[keyName], row[keyIdx])
		}
	}
	out, _ := json.Marshal(m)
	if t.pretty {
		var b bytes.Buffer
		_ = json.Indent(&b, out, "", "\t")
		fmt.Fprintf(t.writer, b.String()+"\n")
	} else {
		fmt.Fprintf(t.writer, string(out))
	}
}

// ColorSupport specifies if the output supports colorized text or not
func (t *JSONShortOutput) ColorSupport() bool {
	// nope. embeds color codes in json values
	return false
}
