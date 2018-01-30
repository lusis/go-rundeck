package cmds

import (
	cli "github.com/lusis/go-rundeck/pkg/cli"
	"github.com/spf13/cobra"
)

func listProjectsFunc(cmd *cobra.Command, args []string) error {
	data, err := cli.Client.ListProjects()
	if err != nil {
		return err
	}

	cli.OutputFormatter.SetHeaders([]string{
		"Name",
		"Description",
		"URL",
	})

	for _, d := range data {
		if rowErr := cli.OutputFormatter.AddRow([]string{
			d.Name,
			d.Description,
			d.URL,
		}); rowErr != nil {
			return rowErr
		}
	}
	cli.OutputFormatter.Draw()
	return nil
}
func listProjectsCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "projects",
		Short: "gets a list of projects from the rundeck server",
		RunE:  listProjectsFunc,
	}
	rootCmd := cli.New(cmd)
	return rootCmd
}
