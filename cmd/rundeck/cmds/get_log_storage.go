package cmds

import (
	"fmt"

	cli "github.com/lusis/go-rundeck/pkg/cli"
	"github.com/spf13/cobra"
)

func getLogStorageFunc(cmd *cobra.Command, args []string) error {
	data, err := cli.Client.GetLogStorage()
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

func getLogStorageCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "get",
		Aliases: []string{"list"},
		Short:   "gets logstorage report from rundeck server",
		RunE:    getLogStorageFunc,
	}
	rootCmd := cli.New(cmd)
	return rootCmd
}
