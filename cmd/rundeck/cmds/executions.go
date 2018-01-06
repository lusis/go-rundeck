package cmds

import "github.com/spf13/cobra"

func executionsCommands() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "executions",
		Short: "operate on rundeck multiple rundeck executions at once",
	}
	cmd.AddCommand(bulkDeleteExecutionsCommand())
	return cmd
}
