package cmds

import (
	"fmt"

	"github.com/lusis/go-rundeck/pkg/cli"
	"github.com/spf13/cobra"
)

func listProjectSCMPluginsFunc(cmd *cobra.Command, args []string) error {
	projectName := args[0]
	data, err := cli.Client.ListSCMPlugins(projectName)
	if err != nil {
		return err
	}
	cli.OutputFormatter.SetHeaders([]string{
		"Title",
		"Description",
		"Enabled?",
		"Configured?",
		"Type",
	})
	for _, p := range data.Import {
		if rowErr := cli.OutputFormatter.AddRow([]string{
			p.Title,
			p.Description,
			fmt.Sprintf("%t", p.Enabled),
			fmt.Sprintf("%t", p.Configured),
			p.Type,
		}); rowErr != nil {
			return err
		}
	}
	for _, p := range data.Export {
		if rowErr := cli.OutputFormatter.AddRow([]string{
			p.Title,
			p.Description,
			fmt.Sprintf("%t", p.Enabled),
			fmt.Sprintf("%t", p.Configured),
			p.Type,
		}); rowErr != nil {
			return err
		}
	}
	cli.OutputFormatter.Draw()
	return nil
}

func listProjectSCMPluginsCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "list-plugins project-name",
		Short: "lists a project's scm plugins",
		Args:  cobra.MinimumNArgs(1),
		RunE:  listProjectSCMPluginsFunc,
	}
	rootCmd := cli.New(cmd)
	return rootCmd
}
