package main

import (
	"errors"
	"fmt"
	"strconv"
	"strings"

	cli "github.com/lusis/go-rundeck/pkg/cli"
	"github.com/spf13/cobra"
)

var (
	projectName string
	maxDelete   int
)

func runFunc(cmd *cobra.Command, args []string) error {
	if projectName == "" {
		return errors.New("project name is required")
	}
	res, err := cli.Client.DeleteAllExecutionsForProject(projectName, maxDelete)
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
		Use:   "rundeck-delete-executions-for-project -p [project-name]",
		Short: "Bulk deletes all executions from a rundeck server for the given project",
		RunE:  runFunc,
	}
	cmd.Flags().StringVarP(&projectName, "project-name", "p", "", "project name")
	cmd.Flags().IntVarP(&maxDelete, "max", "m", 0, "max number of executions to delete")
	rootCmd := cli.New(cmd)
	_ = rootCmd.Execute()
}
