package cli

import (
	"strings"

	rundeck "github.com/lusis/go-rundeck/pkg/rundeck.v21"
	"github.com/lusis/outputter"
	"github.com/spf13/cobra"
)

// Client is a CLI rundeck client instance
var Client *rundeck.Client

// OutputFormat is the output format requested
var OutputFormat string

// OutputFormatter is the configured OutputFormatter
var OutputFormatter outputter.Outputter

// UseFormatter is a flag to specify if we should add the formatting options to the command
var UseFormatter = true

func preRunFunc(cmd *cobra.Command, args []string) error {
	client, err := rundeck.NewClientFromEnv()
	Client = client
	return err
}

// New returns a New rundeck cli object
func New(command *cobra.Command) *cobra.Command {
	command.PreRunE = preRunFunc
	if UseFormatter {
		outputs := outputter.GetOutputters()
		command.PersistentFlags().StringVar(&OutputFormat, "format", "table", "Specify the output format: "+strings.Join(outputs, ","))
		command.PersistentPreRunE = func(cmd *cobra.Command, args []string) error {
			outputFormatter, err := outputter.NewOutputter(OutputFormat)
			if err != nil {
				return err
			}
			OutputFormatter = outputFormatter
			return nil
		}
	}
	return command
}
