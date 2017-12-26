package main

import (
	"errors"
	"fmt"
	"strconv"
	"strings"

	"github.com/lusis/go-rundeck/pkg/cli"
	"github.com/spf13/cobra"
)

var (
	ids = []int{}
)

func runFunc(cmd *cobra.Command, args []string) error {
	if len(ids) == 0 {
		return errors.New("at least one id is required")
	}
	res, err := cli.Client.DeleteExecutions(ids...)
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
		Use:   "rundeck-bulk-delete-executions -i [execution-id][,execution-id]",
		Short: "Bulk deletes executions from a rundeck server",
		RunE:  runFunc,
	}
	cmd.Flags().IntSliceVarP(&ids, "ids", "i", nil, "list of execution ids")
	rootCmd := cli.New(cmd)
	_ = rootCmd.Execute()
}
