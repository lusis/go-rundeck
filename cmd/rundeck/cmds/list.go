package cmds

import (
	"github.com/spf13/cobra"
)

func listProjectExecutionsCommand() *cobra.Command {
	cmd := getProjectExecutionsCommand()
	cmd.Use = "executions project-name [-r] [-m max]"
	return cmd
}

func listCommands() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "list",
		Short: "list various things from the rundeck server",
	}
	cmd.AddCommand(getJobsCommand())
	cmd.AddCommand(listProjectsCommand())
	cmd.AddCommand(listProjectExecutionsCommand())
	return cmd
}
