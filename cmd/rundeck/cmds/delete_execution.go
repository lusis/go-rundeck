package cmds

import (
	cli "github.com/lusis/go-rundeck/pkg/cli"
	"github.com/spf13/cobra"
)

func deleteExecutionFunc(cmd *cobra.Command, args []string) error {
	id := args[0]
	return cli.Client.DeleteExecution(id)
}

func deleteExecutionCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "delete execution-id",
		Short: "deletes an execution on the rundeck server",
		Args:  cobra.MinimumNArgs(1),
		RunE:  deleteExecutionFunc,
	}
	rootCmd := cli.New(cmd)
	rootCmd.ResetFlags()
	return rootCmd
}
