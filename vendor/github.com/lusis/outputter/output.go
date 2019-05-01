package outputter

import (
	"io"
	"sync"
)

var defaultOutput = "tabular"

// OutputRegistry is a registry of all known outputter types
type OutputRegistry struct {
	Outputs map[string]OutputCtr
	sync.RWMutex
}

var registry = &OutputRegistry{Outputs: make(map[string]OutputCtr)}

// OutputCtr returns a func that creates an Outputter
type OutputCtr func() Outputter

// RegisterOutput adds an Outputter to the registry
func RegisterOutput(name string, factory OutputCtr) {
	if _, exists := registry.Outputs[name]; !exists {
		registry.Outputs[name] = factory
	}
}

// Outputter is an interface for multiple ways to draw data to the console
type Outputter interface {
	Draw()
	SetHeaders([]string)
	AddRow([]string) error
	ColorSupport() bool
	SetPretty()
	SetWriter(io.Writer) error
}

// GetDefaultOutputter returns the default output outputters
func GetDefaultOutputter() Outputter {
	f, _ := NewOutputter(defaultOutput)
	return f
}

// GetOutputters returns a list of all supported outputters
func GetOutputters() []string {
	var o []string
	for name := range registry.Outputs {
		o = append(o, name)
	}
	return o
}

// NewOutputter returns a new outputter of the specified type
func NewOutputter(outputFormat string) (Outputter, error) {
	if _, exists := registry.Outputs[outputFormat]; !exists {
		return nil, ErrorUnknownOutputter
	}
	factory, ok := registry.Outputs[outputFormat]
	if !ok {
		return nil, ErrorInvalidOutputter
	}
	o := factory()
	return o, nil
}
