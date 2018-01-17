package cmds

import (
	"github.com/lusis/go-rundeck/pkg/cli"
	"github.com/spf13/cobra"
)

func listJobsFunc(cmd *cobra.Command, args []string) error {
	projectid := args[0]
	data, err := cli.Client.ListJobs(projectid)
	if err != nil {
		return err
	}
	cli.OutputFormatter.SetHeaders([]string{"ID", "Name", "Description", "Group", "Project"})
	for _, d := range data {
		if err := cli.OutputFormatter.AddRow([]string{d.ID, d.Name, d.Description, d.Group, d.Project}); err != nil {
			return err
		}
	}
	cli.OutputFormatter.Draw()
	return nil
}

func getJobsCommand() *cobra.Command {
	getJobsCmd := &cobra.Command{
		Use:   "jobs project-name",
		Short: "lists all jobs for a project",
		RunE:  listJobsFunc,
		Args:  cobra.MinimumNArgs(1),
	}
	cmd := cli.New(getJobsCmd)
	return cmd
}
