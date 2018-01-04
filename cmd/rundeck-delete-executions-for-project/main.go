package main

import (
	"fmt"
	"strconv"
	"strings"

	cli "github.com/lusis/go-rundeck/pkg/cli"
	"github.com/spf13/cobra"
)

var (
	maxDelete int
)

func runFunc(cmd *cobra.Command, args []string) error {
	projectName := args[0]

	eopts := make(map[string]string)
	eopts["max"] = strconv.Itoa(maxDelete)
	e, listErr := cli.Client.ListProjectExecutions(projectName, eopts)
	if listErr != nil {
		return listErr
	}

	var toDelete []int
	for _, execution := range e.Executions {
		toDelete = append(toDelete, execution.ID)
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

func main() {
	cmd := &cobra.Command{
		Use:   "rundeck-delete-executions-for-project project-name [-m X]",
		Short: "Bulk deletes all executions from a rundeck server for the given project",
		Args:  cobra.MinimumNArgs(1),
		RunE:  runFunc,
	}
	cmd.Flags().IntVarP(&maxDelete, "max", "m", 0, "max number of executions to delete")
	rootCmd := cli.New(cmd)
	_ = rootCmd.Execute()
}
