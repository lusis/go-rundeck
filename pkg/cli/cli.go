package cli

import (
	"errors"
	"fmt"
	"strings"

	rundeck "github.com/lusis/go-rundeck/pkg/rundeck"
	"github.com/lusis/outputter"
	"github.com/spf13/cobra"
)

// Client is a CLI rundeck client instance
var Client *rundeck.Client

// OutputFormat is the output format requested
var OutputFormat string

// OutputFormatter is the configured OutputFormatter
var OutputFormatter outputter.Outputter

// TimeFormat is the rundeck time format
var TimeFormat = rundeck.RDTime

func preRunFunc(cmd *cobra.Command, args []string) error {
	client, err := rundeck.NewClientFromEnv()
	Client = client
	return err
}

// BuildParams takes a cobra StringSliceVarP []string and converts it to a map[string]string
func BuildParams(values []string) (map[string]string, error) {
	p := map[string]string{}
	for _, value := range values {
		parts := strings.SplitN(value, "=", 2)
		if len(parts) != 2 {
			return nil, fmt.Errorf("expected key=value got '%s'", value)
		}
		if parts[1] == "" {
			msg := fmt.Sprintf("missing value for key '%s'", parts[0])
			return nil, errors.New(msg)
		}
		if parts[0] == "" {
			msg := fmt.Sprintf("missing key for value '%s'", parts[1])
			return nil, errors.New(msg)
		}
		p[parts[0]] = parts[1]
	}
	return p, nil
}

// New returns a New rundeck cli object
func New(command *cobra.Command) *cobra.Command {
	command.PreRunE = preRunFunc
	command.SilenceUsage = true
	outputs := outputter.GetOutputters()
	command.PersistentFlags().StringVar(&OutputFormat, "output-format", "table", "Specify the output format: "+strings.Join(outputs, ","))
	command.PersistentPreRunE = func(cmd *cobra.Command, args []string) error {
		outputFormatter, err := outputter.NewOutputter(OutputFormat)
		if err != nil {
			return err
		}
		OutputFormatter = outputFormatter
		return nil
	}
	return command
}
