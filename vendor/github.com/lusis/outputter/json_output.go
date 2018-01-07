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
	RegisterOutput("json", newJSONFactoryOutput)
}

// JSONOutput is an Outputter that draws data as json like so:
// [ {"header1":"value1", "header2":"value2"}, {"header1":"value1","header2":"value2"}]
type JSONOutput struct {
	pretty bool
	keys   []string
	writer io.Writer
	data   []map[string]string
	sync.RWMutex
}

func newJSONFactoryOutput() Outputter {
	return NewJSONOutput()
}

// NewJSONOutput creates a new JSONOutput with os.Stdout as the destination
func NewJSONOutput() *JSONOutput {
	return NewJSONOutputWithWriter(os.Stdout)
}

// NewJSONOutputWithWriter returns a new JSONOutput with the specified io.Writer
func NewJSONOutputWithWriter(i io.Writer) *JSONOutput {
	t := &JSONOutput{}
	t.writer = i
	return t
}

// SetHeaders sets the JSON keys to be used
func (t *JSONOutput) SetHeaders(headers []string) {
	t.keys = headers
}

// AddRow adds a new set of data to the JSON array
func (t *JSONOutput) AddRow(row []string) error {
	if len(t.keys) == 0 {
		return ErrorOutputAddRowNoHeaders
	}
	if len(t.keys) < len(row) {
		return ErrorOutputAddRowTooFewHeaders
	}
	t.Lock()
	defer t.Unlock()
	m := make(map[string]string)
	if len(row) < len(t.keys) {
		difference := len(t.keys) - len(row)
		// we have to account for this and fill in empty values
		missingVals := make([]string, difference)
		row = append(row, missingVals...)
	}
	for keyIdx, keyName := range t.keys {
		m[keyName] = row[keyIdx]
	}
	t.data = append(t.data, m)
	return nil
}

// SetPretty sets json output to pretty format
func (t *JSONOutput) SetPretty() {
	t.Lock()
	defer t.Unlock()
	t.pretty = true
}

// SetWriter sets the output io.Writer
func (t *JSONOutput) SetWriter(i io.Writer) error {
	t.Lock()
	defer t.Unlock()
	t.writer = i
	return nil
}

// Draw renders the table to the configured io.Writer
func (t *JSONOutput) Draw() {
	t.RLock()
	defer t.RUnlock()

	out, _ := json.Marshal(t.data)
	if t.pretty {
		var b bytes.Buffer
		_ = json.Indent(&b, out, "", "\t")
		fmt.Fprintf(t.writer, b.String()+"\n")
	} else {
		fmt.Fprintf(t.writer, string(out))
	}
}

// ColorSupport specifies if the output supports colorized text or not
func (t *JSONOutput) ColorSupport() bool {
	// nope. embeds color codes in json values
	return false
}
