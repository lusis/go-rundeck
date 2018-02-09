package cmds

import (
	"fmt"
	"strconv"

	cli "github.com/lusis/go-rundeck/pkg/cli"
	"github.com/spf13/cobra"
)

func getExecutionOutputFunc(cmd *cobra.Command, args []string) error {
	id := args[0]
	eID, eIDerr := strconv.Atoi(id)
	if eIDerr != nil {
		return eIDerr
	}
	data, err := cli.Client.GetExecutionOutput(eID)
	if err != nil {
		return err
	}
	for _, entry := range data.Entries {
		fmt.Println(entry.Log)
	}
	return nil
}

func getExecutionOutputCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "output execution-id",
		Short: "gets an execution's output from the rundeck server",
		Args:  cobra.MinimumNArgs(1),
		RunE:  getExecutionOutputFunc,
	}
	rootCmd := cli.New(cmd)
	rootCmd.ResetFlags()
	return rootCmd
}
