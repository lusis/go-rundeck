package cmds

import (
	"fmt"
	"strings"

	"github.com/lusis/go-rundeck/pkg/cli"
	"github.com/spf13/cobra"
)

var (
	createProjectOptions []string
	createProjectCmd     = &cobra.Command{
		Use:   "create project-name [-o foo=bar]",
		Short: "creates a rundeck project",
		RunE:  createProjectFunc,
		Args:  cobra.MinimumNArgs(1),
	}
)

func createProjectFunc(cmd *cobra.Command, args []string) error {
	projectName := args[0]
	params, paramErr := cli.BuildParams(createProjectOptions)
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

func createProjectCommand() *cobra.Command {
	cmd := cli.New(createProjectCmd)
	cmd.Flags().StringSliceVarP(&createProjectOptions, "options", "o", []string{}, "custom options to pass in format of name=value. Can specify multiple times")
	return cmd
}
