package cmds

import (
	"github.com/lusis/go-rundeck/pkg/cli"
	"github.com/spf13/cobra"
)

func getProjectConfigFunc(cmd *cobra.Command, args []string) error {
	projectName := args[0]
	data, err := cli.Client.GetProjectConfiguration(projectName)
	if err != nil {
		return err
	}
	cli.OutputFormatter.SetHeaders([]string{
		"Name",
		"Value",
	})
	for k, v := range data {
		if err := cli.OutputFormatter.AddRow([]string{
			k,
			v}); err != nil {
			return err
		}
	}
	cli.OutputFormatter.Draw()
	return nil
}

func getProjectConfigCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "config project-name",
		Short: "gets a project's configuration from a rundeck server",
		Args:  cobra.MinimumNArgs(1),
		RunE:  getProjectConfigFunc,
	}
	rootCmd := cli.New(cmd)
	return rootCmd
}
