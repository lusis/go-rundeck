package cmds

import (
	"fmt"
	"os"

	"github.com/lusis/go-rundeck/pkg/cli"
	rundeck "github.com/lusis/go-rundeck/pkg/rundeck"
	"github.com/spf13/cobra"
)

func runAdHocScriptFunc(cmd *cobra.Command, args []string) error {
	projectID := args[0]
	script := args[1]
	options := []rundeck.AdHocScriptOption{
		rundeck.ScriptThreadCount(adHocNodeThreadCount),
		rundeck.ScriptKeepGoing(adHocNodeKeepGoing),
		rundeck.ScriptArgsQuoted(adHocArgsQuoted),
		rundeck.ScriptInterpreter(adHocScriptInterpreter),
	}
	if &adHocAsUser != nil {
		options = append(options, rundeck.ScriptRunAs(adHocAsUser))
	}
	if &adHocArgString != nil {
		options = append(options, rundeck.ScriptArgString(adHocArgString))
	}
	if &adHocFileExtension != nil {
		options = append(options, rundeck.ScriptFileExtension(adHocFileExtension))
	}
	if &adHocFilter != nil {
		options = append(options, rundeck.ScriptNodeFilters(adHocFilter)) // nolint: ineffassign
	}
	scriptData, scriptErr := os.Open(script)
	if scriptErr != nil {
		return scriptErr
	}
	res, err := cli.Client.RunAdHocScript(projectID, scriptData, options...)
	if err != nil {
		return err
	}
	data, dataErr := cli.Client.GetExecutionInfo(res.Execution.ID)
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

func runAdHocScriptCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "script project-name script-file",
		Short: "uploads the specified script and runs it against a project on the rundeck server",
		Args:  cobra.MinimumNArgs(2),
		RunE:  runAdHocScriptFunc,
	}
	rootCmd := cli.New(cmd)
	rootCmd.Flags().StringVar(&adHocScriptInterpreter, "interpreter", "/bin/bash", "the script interpreter to use")
	rootCmd.Flags().StringVar(&adHocFileExtension, "file-extension", "", "file extension to use on the remote node")
	rootCmd.Flags().StringVar(&adHocArgString, "argstring", "", "args string to pass to the script")
	rootCmd.Flags().BoolVar(&adHocArgsQuoted, "args-quoted", false, "should arguments to interpreter be quoted")
	return rootCmd
}
