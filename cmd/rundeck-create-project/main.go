package main

import (
	"errors"
	"fmt"
	"strings"

	cli "github.com/lusis/go-rundeck/pkg/cli"
	"github.com/spf13/cobra"
)

var (
	options []string
)

func buildParams(values []string) (map[string]string, error) {
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

func runFunc(cmd *cobra.Command, args []string) error {
	projectName := args[0]
	params, paramErr := buildParams(options)
	if paramErr != nil {
		return paramErr
	}
	data, err := cli.Client.CreateProject(projectName, params)
	if err != nil {
		return err
	}
	cli.OutputFormatter.SetHeaders([]string{
		"URL",
		"Name",
		"Description",
		"Config",
	})
	var config []string
	for k, v := range data.Properties {
		config = append(config, fmt.Sprintf("%s - %s ", k, v))
	}
	if rowErr := cli.OutputFormatter.AddRow([]string{
		data.URL,
		data.Name,
		data.Description,
		strings.Join(config, "\n"),
	}); rowErr != nil {
		return rowErr
	}
	cli.OutputFormatter.Draw()
	return nil
}
func main() {
	cmd := &cobra.Command{
		Use:   "rundeck-create-project project-name [-o foo=bar]",
		Short: "creates a rundeck project",
		RunE:  runFunc,
		Args:  cobra.MinimumNArgs(1),
	}
	cmd.Flags().StringSliceVarP(&options, "options", "o", []string{}, "custom options to pass in format of name=value. Can specify multiple times")
	rootCmd := cli.New(cmd)
	_ = rootCmd.Execute()
}
