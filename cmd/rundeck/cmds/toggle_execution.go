package cmds

import (
	"errors"
	"strconv"

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
	i, iErr := strconv.Atoi(id)
	if iErr != nil {
		return iErr
	}
	_, err := cli.Client.DisableExecution(i)
	return err

}

func enableExecutionFunc(cmd *cobra.Command, args []string) error {
	id := args[0]
	i, iErr := strconv.Atoi(id)
	if iErr != nil {
		return iErr
	}
	_, err := cli.Client.EnableExecution(i)
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
