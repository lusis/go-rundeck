package cmds

import (
	"fmt"
	"strings"

	cli "github.com/lusis/go-rundeck/pkg/cli"
	"github.com/spf13/cobra"
)

var (
	getLogStorageIncompleteOnly bool
)

func getLogStorageFunc(cmd *cobra.Command, args []string) error {
	data, err := cli.Client.GetLogStorageInfo()
	if err != nil {
		return err
	}
	headers := []string{
		"Enabled?",
		"Plugin Name",
		"Succeeded",
		"Failed",
		"Queued",
		"Total",
		"Incomplete",
		"Missing",
	}
	cli.OutputFormatter.SetHeaders(headers)
	if rowErr := cli.OutputFormatter.AddRow([]string{
		fmt.Sprintf("%t", data.Enabled),
		data.PluginName,
		fmt.Sprintf("%d", data.SucceededCount),
		fmt.Sprintf("%d", data.FailedCount),
		fmt.Sprintf("%d", data.QueuedCount),
		fmt.Sprintf("%d", data.TotalCount),
		fmt.Sprintf("%d", data.IncompleteCount),
		fmt.Sprintf("%d", data.MissingCount),
	}); rowErr != nil {
		return rowErr
	}
	cli.OutputFormatter.Draw()
	return nil
}

func getIncompleteLogStorageFunc(cmd *cobra.Command, args []string) error {
	data, err := cli.Client.GetIncompleteLogStorage()
	if err != nil {
		return err
	}
	headers := []string{
		"ID",
		"Project",
		"HRef",
		"Permalink",
		"Local Files Present?",
		"Incomplete File Types",
		"Queued?",
		"Failed?",
		"Date",
		"Errors",
	}
	cli.OutputFormatter.SetHeaders(headers)
	for _, e := range data.Executions {
		if rowErr := cli.OutputFormatter.AddRow([]string{
			fmt.Sprintf("%d", e.ID),
			e.Project,
			e.HRef,
			e.Permalink,
			fmt.Sprintf("%t", e.Storage.LocalFilesPresent),
			e.Storage.IncompleteFiletypes,
			fmt.Sprintf("%t", e.Storage.Queued),
			fmt.Sprintf("%t", e.Storage.Failed),
			e.Storage.Date.String(),
			strings.Join(e.Errors, "\n"),
		}); rowErr != nil {
			return rowErr
		}
	}
	cli.OutputFormatter.Draw()
	return nil
}

func getLogStorageWrapper(cmd *cobra.Command, args []string) error {
	if getLogStorageIncompleteOnly {
		return getIncompleteLogStorageFunc(cmd, args)
	}
	return getLogStorageFunc(cmd, args)
}

func getLogStorageCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "get [--incomplete]",
		Aliases: []string{"list"},
		Short:   "gets logstorage report from rundeck server",
		RunE:    getLogStorageWrapper,
	}
	rootCmd := cli.New(cmd)
	rootCmd.Flags().BoolVar(&getLogStorageIncompleteOnly, "incomplete", false, "return executions with incomplete logstorage")
	return rootCmd
}
