package cmds

import (
	"fmt"

	"github.com/lusis/go-rundeck/pkg/cli"
	"github.com/spf13/cobra"
)

func getJobFunc(cmd *cobra.Command, args []string) error {
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
	description := ""
	if data.Description != "" {
		description = data.Description
	}
	group := "<none>"
	if data.Group != "" {
		group = data.Group
	}
	if rowErr := cli.OutputFormatter.AddRow([]string{
		data.ID,
		data.Name,
		description,
		group,
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

func getJobCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "get job-id",
		Short: "gets job metadata from a rundeck server",
		Args:  cobra.MinimumNArgs(1),
		RunE:  getJobFunc,
	}
	rootCmd := cli.New(cmd)
	return rootCmd
}
