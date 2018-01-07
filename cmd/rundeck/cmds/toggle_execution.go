package cmds

import (
	"errors"

	"github.com/lusis/go-rundeck/pkg/cli"
	"github.com/spf13/cobra"
)

func toggleExecutionFunc(cmd *cobra.Command, args []string) error {
	action := args[1]
	switch action {
	case enable:
		return enableExecutionFunc(cmd, args)
	case disable:
		return disableExecutionFunc(cmd, args)
	default:
		return errors.New("action must be one of 'enable' or 'disable'")
	}
}

func disableExecutionFunc(cmd *cobra.Command, args []string) error {
	id := args[0]
	_, err := cli.Client.DisableExecution(id)
	return err

}

func enableExecutionFunc(cmd *cobra.Command, args []string) error {
	id := args[0]
	_, err := cli.Client.EnableExecution(id)
	return err

}

func toggleExecutionCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "toggle job-id [enable|disable]",
		Short: "enables or disables an execution on the rundeck server",
		Args:  cobra.MinimumNArgs(2),
		RunE:  toggleExecutionFunc,
	}
	rootCmd := cli.New(cmd)
	rootCmd.ResetFlags()
	return rootCmd
}
