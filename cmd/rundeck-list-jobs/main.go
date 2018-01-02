package main

import (
	cli "github.com/lusis/go-rundeck/pkg/cli"
	"github.com/spf13/cobra"
)

func runFunc(cmd *cobra.Command, args []string) error {
	projectid := args[0]
	data, err := cli.Client.ListJobs(projectid)
	if err != nil {
		return err
	}
	cli.OutputFormatter.SetHeaders([]string{"ID", "Name", "Description", "Group", "Project"})
	for _, d := range *data {
		if err := cli.OutputFormatter.AddRow([]string{d.ID, d.Name, d.Description, d.Group, d.Project}); err != nil {
			return err
		}
	}
	cli.OutputFormatter.Draw()
	return nil

}
func main() {
	cmd := &cobra.Command{
		Use:   "rundeck-list-jobs project-name",
		Short: "gets a list of jobs for a project from the rundeck server",
		Args:  cobra.MinimumNArgs(1),
		RunE:  runFunc,
	}
	rootCmd := cli.New(cmd)
	_ = rootCmd.Execute()
}
