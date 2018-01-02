package main

import (
	"errors"
	"fmt"

	cli "github.com/lusis/go-rundeck/pkg/cli"
	"github.com/spf13/cobra"
)

var (
	projectid string
	jobname   string
)

func runFunc(cmd *cobra.Command, args []string) error {
	if projectid == "" || jobname == "" {
		return errors.New("You must specify a project name and job name")
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
		Use:   "rundeck-find-job-by-name -p [project-name] -j [job-name]",
		Short: "finds a project's job by name",
		RunE:  runFunc,
	}
	cmd.Flags().StringVarP(&projectid, "project-id", "p", "", "project id")
	cmd.Flags().StringVarP(&jobname, "job-name", "j", "", "job name")
	rootCmd := cli.New(cmd)
	_ = rootCmd.Execute()
}
