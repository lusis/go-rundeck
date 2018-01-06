package cmds

import "github.com/spf13/cobra"

func jobCommands() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "job",
		Short: "operate on individual rundeck jobs",
	}
	cmd.AddCommand(runJobCommand())
	cmd.AddCommand(deleteJobCommand())
	cmd.AddCommand(getJobCommand())
	cmd.AddCommand(getJobOptsCommand())
	cmd.AddCommand(exportJobCommand())
	cmd.AddCommand(importJobCommand())
	cmd.AddCommand(findJobByNameCommand())
	return cmd
}
