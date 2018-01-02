package main

import (
	"fmt"

	"github.com/lusis/go-rundeck/pkg/cli"
	"github.com/spf13/cobra"
)

func runFunc(cmd *cobra.Command, args []string) error {
	jobid := args[0]
	data, err := cli.Client.GetJobMetaData(jobid)
	if err != nil {
		return err
	}
	cli.OutputFormatter.SetHeaders([]string{
		"ID",
		"Name",
		"Description",
		"Group",
		"Scheduled?",
		"Schedule Enabled?",
		"Enabled?",
		"Average Duration",
	})
	if rowErr := cli.OutputFormatter.AddRow([]string{
		data.ID,
		data.Description,
		data.Group,
		fmt.Sprintf("%t", data.Scheduled),
		fmt.Sprintf("%t", data.ScheduleEnabled),
		fmt.Sprintf("%t", data.Enabled),
		fmt.Sprintf("%d", data.AverageDuration),
	}); rowErr != nil {
		return rowErr
	}
	cli.OutputFormatter.Draw()
	return nil
}

func main() {
	cmd := &cobra.Command{
		Use:   "rundeck-get-job job-id",
		Short: "gets job metadata from a rundeck server",
		Args:  cobra.MinimumNArgs(1),
		RunE:  runFunc,
	}
	rootCmd := cli.New(cmd)
	_ = rootCmd.Execute()
}
