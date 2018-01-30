package cmds

import (
	"fmt"
	"strconv"

	cli "github.com/lusis/go-rundeck/pkg/cli"
	"github.com/spf13/cobra"
)

func getExecutionFunc(cmd *cobra.Command, args []string) error {
	id := args[0]
	eID, eIDerr := strconv.Atoi(id)
	if eIDerr != nil {
		return eIDerr
	}
	data, err := cli.Client.GetExecutionInfo(eID)
	if err != nil {
		return err
	}
	cli.OutputFormatter.SetHeaders([]string{
		"ID",
		"User",
		"Status",
		"Start Date",
		"End Date",
		"Args",
		"Server UUID",
		"Custom Status",
	})
	var dateEnded string
	if data.DateEnded.Date == nil {
		dateEnded = "running" // nolint: goconst
	} else {
		dateEnded = data.DateEnded.Date.String()
	}
	if rowErr := cli.OutputFormatter.AddRow([]string{
		fmt.Sprintf("%d", data.ID),
		data.User,
		data.Status,
		data.DateStarted.Date.String(),
		dateEnded,
		data.ArgString,
		data.ServerUUID,
		data.CustomStatus,
	}); rowErr != nil {
		return rowErr
	}

	cli.OutputFormatter.Draw()
	return nil
}
func getExecutionCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "get execution-id",
		Short: "gets an execution from the rundeck server",
		Args:  cobra.MinimumNArgs(1),
		RunE:  getExecutionFunc,
	}
	rootCmd := cli.New(cmd)
	return rootCmd
}
