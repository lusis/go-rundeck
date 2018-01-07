package cmds

import (
	"fmt"
	"strings"

	"github.com/lusis/go-rundeck/pkg/cli"
	"github.com/spf13/cobra"
)

func getProjectFunc(cmd *cobra.Command, args []string) error {
	projectName := args[0]
	data, err := cli.Client.GetProjectInfo(projectName)
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

func getProjectCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "get project-name",
		Short: "gets project info from the rundeck server",
		Args:  cobra.MinimumNArgs(1),
		RunE:  getProjectFunc,
	}
	rootCmd := cli.New(cmd)
	return rootCmd
}
