package main

import (
	"errors"
	"fmt"

	cli "github.com/lusis/go-rundeck/pkg/cli"
	"github.com/spf13/cobra"
)

var (
	projectid string
)

func runFunc(cmd *cobra.Command, args []string) error {
	jobname := args[0]
	if projectid == "" {
		return errors.New("You must specify a project name")
	}
	data, err := cli.Client.FindJobByName(jobname, projectid)
	if err != nil {
		fmt.Printf("%s\n", err)
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
		Use:   "rundeck-find-job-by-name job-name -p project-name",
		Short: "finds a project's job by name",
		Args:  cobra.MinimumNArgs(1),
		RunE:  runFunc,
	}
	cmd.Flags().StringVarP(&projectid, "project-id", "p", "", "project id")
	rootCmd := cli.New(cmd)
	_ = rootCmd.Execute()
}
