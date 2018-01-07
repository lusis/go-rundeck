package cmds

import "github.com/spf13/cobra"

func jobsCommands() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "jobs",
		Short: "operate on rundeck multiple rundeck jobs at once",
	}
	cmd.AddCommand(bulkToggleExecutionsCommand())
	cmd.AddCommand(bulkToggleScheduleCommand())
	return cmd
}
