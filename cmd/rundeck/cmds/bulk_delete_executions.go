package cmds

import (
	"errors"
	"fmt"
	"strconv"
	"strings"

	"github.com/lusis/go-rundeck/pkg/cli"
	"github.com/spf13/cobra"
)

var (
	bulkDeleteExecutionIDs = []int{}
)

func bulkDeleteExecutionsFunc(cmd *cobra.Command, args []string) error {
	if len(bulkDeleteExecutionIDs) == 0 {
		return errors.New("at least one id is required")
	}
	res, err := cli.Client.DeleteExecutions(bulkDeleteExecutionIDs...)
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

func bulkDeleteExecutionsCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "delete -i [execution-id][,execution-id]",
		Short: "Bulk deletes executions from a rundeck server",
		RunE:  bulkDeleteExecutionsFunc,
	}
	rootCmd := cli.New(cmd)
	rootCmd.Flags().IntSliceVarP(&bulkDeleteExecutionIDs, "ids", "i", nil, "list of execution ids")

	return rootCmd
}
