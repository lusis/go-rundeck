package main

import (
	"errors"
	"fmt"

	"github.com/lusis/go-rundeck/pkg/cli"
	"github.com/spf13/cobra"
)

var (
	jobid string
)

func runFunc(cmd *cobra.Command, args []string) error {
	if jobid == "" {
		return errors.New("job id is required")
	}
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
		Use:   "rundeck-get-job -j [job-id]",
		Short: "gets job metadata from a rundeck server",
		RunE:  runFunc,
	}
	cmd.Flags().StringVarP(&jobid, "job-id", "j", "", "job id")
	rootCmd := cli.New(cmd)
	_ = rootCmd.Execute()
}
