package main

import (
	"errors"
	"strconv"

	cli "github.com/lusis/go-rundeck/pkg/cli"
	"github.com/spf13/cobra"
)

var (
	projectid string
)

func runFunc(cmd *cobra.Command, args []string) error {
	if projectid == "" {
		return errors.New("you must specify a project id")
	}
	data, err := cli.Client.ListRunningExecutions(projectid)
	if err != nil {
		return err
	}
	cli.OutputFormatter.SetHeaders([]string{
		"ID",
		"Job Name",
		"Job Description",
		"Arguments",
		"Node Success/Failure Count",
		"User",
		"Start",
		"Project",
	})
	for _, d := range data.Executions {
		var description string
		var name string
		if &d.Job != nil {
			name = d.Job.Name
			description = d.Job.Description
		} else {
			name = "<adhoc>"
			description = d.Description
		}
		if rowErr := cli.OutputFormatter.AddRow([]string{
			strconv.Itoa(d.ID),
			name,
			description,
			d.ArgString,
			strconv.Itoa(len(d.SuccessfulNodes)) + "/" + strconv.Itoa(len(d.FailedNodes)),
			d.User,
			d.DateStarted.Date.String(),
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
		Use:   "rundeck-list-running-executions -p project-id",
		Short: "gets a list of running executions for a project from the rundeck server",
		RunE:  runFunc,
	}
	cmd.Flags().StringVarP(&projectid, "project-id", "p", "", "project id")
	rootCmd := cli.New(cmd)
	_ = rootCmd.Execute()
}
