package main

import (
	"log"

	outputter "github.com/lusis/outputter"
	kingpin "gopkg.in/alecthomas/kingpin.v2"
)

func main() {
	var outputFormat string
	kingpin.Flag("format", "format for output").
		Default("tabular").
		EnumVar(&outputFormat, outputter.GetOutputters()...)
	kingpin.Parse()
	outputFormatter, err := outputter.NewOutputter(outputFormat)
	if err != nil {
		log.Fatalf("unable to create an outputter: %s", err.Error())
	}

	outputFormatter.SetHeaders([]string{"header1", "header2", "header3"})
	rowErr := outputFormatter.AddRow([]string{"value1", "value2", "value3"})
	if rowErr != nil {
		log.Fatalf("error adding row: %s", rowErr.Error())
	}
	outputFormatter.Draw()
}
