package cmds

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/lusis/go-rundeck/pkg/cli"
	"github.com/spf13/cobra"
)

var getProjectExecutionsMax string
var getProjectExecutionsRunningOnly bool
var deleteProjectExecutionsMax int

func deleteProjectExecutionsFunc(cmd *cobra.Command, args []string) error {
	projectName := args[0]

	eopts := make(map[string]string)
	eopts["max"] = strconv.Itoa(deleteProjectExecutionsMax)
	e, listErr := cli.Client.ListProjectExecutions(projectName, eopts)
	if listErr != nil {
		return listErr
	}

	var toDelete []int
	for _, execution := range e.Executions {
		toDelete = append(toDelete, execution.ID)
	}
	if len(toDelete) == 0 {
		fmt.Println("no executions to delete")
		return nil
	}
	res, err := cli.Client.DeleteExecutions(toDelete...)
	if err != nil {
		return err
	}
	cli.OutputFormatter.SetHeaders([]string{
		"Successful",
		"Failed",
		"Errors",
	})
	var failureResults []string
	if res.FailedCount != 0 {
		for _, f := range res.Failures {
			failureResults = append(failureResults, fmt.Sprintf("%s: %s", f.ID, f.Message))
		}
	}
	if rowErr := cli.OutputFormatter.AddRow([]string{
		strconv.Itoa(res.SuccessCount),
		strconv.Itoa(res.FailedCount),
		strings.Join(failureResults, "\n"),
	}); rowErr != nil {
		return rowErr
	}
	cli.OutputFormatter.Draw()
	return nil
}

func getProjectExecutionsWrapperFunc(cmd *cobra.Command, args []string) error {
	if getProjectExecutionsRunningOnly {
		return getRunningProjectExectionsFunc(cmd, args)
	}
	return getProjectExecutionsFunc(cmd, args)

}
func getProjectExecutionsFunc(cmd *cobra.Command, args []string) error {
	projectid := args[0]
	options := make(map[string]string)
	options["max"] = getProjectExecutionsMax
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
			name = adhoc
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

func getRunningProjectExectionsFunc(cmd *cobra.Command, args []string) error {
	projectid := args[0]
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
			name = adhoc
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

func getProjectExecutionsCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "list project-name [-r] [-m max]",
		Short: "gets a list of executions for a project from the rundeck server optionally only running executions",
		Args:  cobra.MinimumNArgs(1),
		RunE:  getProjectExecutionsWrapperFunc,
	}
	rootCmd := cli.New(cmd)
	rootCmd.Flags().StringVarP(&getProjectExecutionsMax, "max", "m", "", "max results")
	rootCmd.Flags().BoolVarP(&getProjectExecutionsRunningOnly, "running-only", "r", false, "show only running executions")
	return rootCmd
}

func deleteProjectExecutionsCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "delete project-name [-m X]",
		Short: "Bulk deletes all executions from a rundeck server for the given project",
		Args:  cobra.MinimumNArgs(1),
		RunE:  deleteProjectExecutionsFunc,
	}
	cmd.Flags().IntVarP(&deleteProjectExecutionsMax, "max", "m", 0, "max number of executions to delete")
	rootCmd := cli.New(cmd)
	return rootCmd
}
func projectExecutionsCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "executions",
		Short: "operates on a project's executions",
	}
	cmd.AddCommand(getProjectExecutionsCommand())
	cmd.AddCommand(deleteProjectExecutionsCommand())
	return cmd
}
