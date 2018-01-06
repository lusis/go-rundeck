package cmds

import (
	"github.com/lusis/go-rundeck/pkg/cli"
	"github.com/spf13/cobra"
)

func deleteJobFunc(cmd *cobra.Command, args []string) error {
	id := args[0]
	return cli.Client.DeleteJob(id)
}

func deleteJobCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "delete job-id",
		Short: "deletes a job on the rundeck server",
		Args:  cobra.MinimumNArgs(1),
		RunE:  deleteJobFunc,
	}
	rootCmd := cli.New(cmd)
	rootCmd.ResetFlags()
	return rootCmd
}
