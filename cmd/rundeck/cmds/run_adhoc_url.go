package cmds

import (
	"fmt"

	"github.com/lusis/go-rundeck/pkg/cli"
	rundeck "github.com/lusis/go-rundeck/pkg/rundeck"
	"github.com/spf13/cobra"
)

func runAdHocURLFunc(cmd *cobra.Command, args []string) error {
	projectID := args[0]
	script := args[1]
	options := []rundeck.AdHocScriptURLOption{
		rundeck.ScriptURLThreadCount(adHocNodeThreadCount),
		rundeck.ScriptURLKeepGoing(adHocNodeKeepGoing),
		rundeck.ScriptURLArgsQuoted(adHocArgsQuoted),
		rundeck.ScriptURLInterpreter(adHocScriptInterpreter),
	}
	if len(adHocAsUser) > 0 {
		options = append(options, rundeck.ScriptURLRunAs(adHocAsUser))
	}
	if len(adHocArgString) > 0 {
		options = append(options, rundeck.ScriptURLArgString(adHocArgString))
	}
	if len(adHocFileExtension) > 0 {
		options = append(options, rundeck.ScriptURLFileExtension(adHocFileExtension))
	}
	if len(adHocFilter) > 0 {
		options = append(options, rundeck.ScriptURLNodeFilters(adHocFilter)) // nolint: ineffassign
	}

	res, err := cli.Client.RunAdHocScriptFromURL(projectID, script, options...)
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

func runAdHocURLCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "url project-name script-url",
		Short: "runs script at specified url against a project on the rundeck server",
		Args:  cobra.MinimumNArgs(2),
		RunE:  runAdHocURLFunc,
	}
	rootCmd := cli.New(cmd)
	rootCmd.Flags().StringVar(&adHocScriptInterpreter, "interpreter", "/bin/bash", "the script interpreter to use")
	rootCmd.Flags().StringVar(&adHocFileExtension, "file-extension", "", "file extension to use on the remote node")
	rootCmd.Flags().StringVar(&adHocArgString, "argstring", "", "args string to pass to the script")
	rootCmd.Flags().BoolVar(&adHocArgsQuoted, "args-quoted", false, "should arguments to interpreter be quoted")
	return rootCmd
}
