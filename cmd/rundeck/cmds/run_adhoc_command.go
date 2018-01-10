package cmds

import (
	"fmt"
	"strings"

	"github.com/lusis/go-rundeck/pkg/cli"
	rundeck "github.com/lusis/go-rundeck/pkg/rundeck"
	"github.com/spf13/cobra"
)

func runAdHocCmdFunc(cmd *cobra.Command, args []string) error {
	projectID := args[0]
	command := args[1:]
	options := []rundeck.AdHocRunOption{
		rundeck.CmdThreadCount(adHocNodeThreadCount),
		rundeck.CmdKeepGoing(adHocNodeKeepGoing),
	}
	if &adHocAsUser != nil {
		options = append(options, rundeck.CmdRunAs(adHocAsUser))
	}
	if &adHocFilter != nil {
		options = append(options, rundeck.CmdNodeFilters(adHocFilter)) // nolint: ineffassign
	}
	res, err := cli.Client.RunAdHocCommand(projectID, strings.Join(command, " "), options...)
	if err != nil {
		return err
	}
	data, dataErr := cli.Client.GetExecutionInfo(fmt.Sprintf("%d", res.Execution.ID))
	if dataErr != nil {
		return fmt.Errorf("Execution started but unable to get information: %s", dataErr.Error())
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
		dateEnded = ""
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

func runAdHocCmdCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "run project-name command",
		Short: "runs the specified command against a project on the rundeck server. Consider quoting command",
		Args:  cobra.MinimumNArgs(2),
		RunE:  runAdHocCmdFunc,
	}
	rootCmd := cli.New(cmd)
	return rootCmd
}
