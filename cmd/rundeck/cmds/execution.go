package cmds

import "github.com/spf13/cobra"

func executionCommands() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "execution",
		Short: "operate on individual rundeck executions",
	}
	cmd.AddCommand(getExecutionCommand())
	cmd.AddCommand(deleteExecutionCommand())
	cmd.AddCommand(toggleExecutionCommand())
	cmd.AddCommand(getExecutionOutputCommand())
	return cmd
}
