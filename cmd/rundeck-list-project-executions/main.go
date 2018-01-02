package main

import (
	"strconv"

	cli "github.com/lusis/go-rundeck/pkg/cli"
	"github.com/spf13/cobra"
)

var (
	max string
)

func runFunc(cmd *cobra.Command, args []string) error {
	projectid := args[0]
	options := make(map[string]string)
	options["max"] = max
	data, err := cli.Client.ListProjectExecutions(projectid, options)
	if err != nil {
		return err
	}
	cli.OutputFormatter.SetHeaders([]string{
		"ID",
		"Job Name",
		"Job Description",
		"Status",
		"Node Success/Failure Count",
		"User",
		"Start",
		"End",
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
		dateEnded := ""
		if d.DateEnded.Date != nil {
			dateEnded = d.DateEnded.Date.String()
		}
		if rowErr := cli.OutputFormatter.AddRow([]string{
			strconv.Itoa(d.ID),
			name,
			description,
			d.Status,
			strconv.Itoa(len(d.SuccessfulNodes)) + "/" + strconv.Itoa(len(d.FailedNodes)),
			d.User,
			d.DateStarted.Date.String(),
			dateEnded,
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
		Use:   "rundeck-list-project-executions project-name [-m max]",
		Short: "gets a list of executions for a project from the rundeck server",
		Args:  cobra.MinimumNArgs(1),
		RunE:  runFunc,
	}
	cmd.Flags().StringVarP(&max, "max", "m", "", "max results")
	rootCmd := cli.New(cmd)
	_ = rootCmd.Execute()

}
