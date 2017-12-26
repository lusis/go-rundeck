package main

import (
	"errors"
	"fmt"

	"github.com/lusis/go-rundeck/pkg/cli"
	"github.com/spf13/cobra"
)

var (
	projectid string
)

func runFunc(cmd *cobra.Command, args []string) error {
	if projectid == "" {
		return errors.New("project id is required")
	}
	data, err := cli.Client.GetHistory(projectid)
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
func main() {
	cmd := &cobra.Command{
		Use:   "rundeck-get-history -p [project-id]",
		Short: "gets project history from the rundeck server",
		RunE:  runFunc,
	}
	cmd.Flags().StringVarP(&projectid, "project-id", "p", "", "project id")
	rootCmd := cli.New(cmd)
	_ = rootCmd.Execute()
}
