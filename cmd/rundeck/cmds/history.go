package cmds

import (
	"fmt"

	"github.com/lusis/go-rundeck/pkg/cli"
	"github.com/spf13/cobra"
)

func projectHistoryFunc(cmd *cobra.Command, args []string) error {
	projectid := args[0]
	data, err := cli.Client.ListHistory(projectid)
	if err != nil {
		return err
	}
	cli.OutputFormatter.SetHeaders([]string{
		"Title",
		"Status",
		"Summary",
		"Start Time",
		"End Time",
		"S/F/T",
		"Execution",
		"User",
		"Project",
	})
	for _, d := range data.Events {
		if rowErr := cli.OutputFormatter.AddRow([]string{
			d.Title,
			d.Status,
			d.Summary,
			d.DateStarted.String(),
			d.DateEnded.String(),
			fmt.Sprintf("%d/%d/%d", d.NodeSummary.Succeeded, d.NodeSummary.Failed, d.NodeSummary.Total),
			d.Execution.ID,
			d.User,
			d.Project,
		}); rowErr != nil {
			return rowErr
		}
	}

	cli.OutputFormatter.Draw()
	return nil
}
func projectHistoryCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "history project-name",
		Short: "gets project history from the rundeck server",
		Args:  cobra.MinimumNArgs(1),
		RunE:  projectHistoryFunc,
	}
	rootCmd := cli.New(cmd)
	return rootCmd
}
