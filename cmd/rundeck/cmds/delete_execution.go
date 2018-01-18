package cmds

import (
	"strconv"

	cli "github.com/lusis/go-rundeck/pkg/cli"
	"github.com/spf13/cobra"
)

func deleteExecutionFunc(cmd *cobra.Command, args []string) error {
	id := args[0]
	eID, err := strconv.Atoi(id)
	if err != nil {
		return err
	}
	return cli.Client.DeleteExecution(eID)
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
