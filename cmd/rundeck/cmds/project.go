package cmds

import (
	"github.com/spf13/cobra"
)

func projectCommands() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "project",
		Short: "operate on a rundeck project",
	}
	cmd.AddCommand(getProjectCommand())
	cmd.AddCommand(deleteProjectCommand())
	cmd.AddCommand(createProjectCommand())
	cmd.AddCommand(getJobsCommand())
	cmd.AddCommand(projectExecutionsCommand())
	cmd.AddCommand(projectHistoryCommand())
	cmd.AddCommand(getProjectConfigCommand())
	cmd.AddCommand(exportProjectCommand())
	return cmd
}
