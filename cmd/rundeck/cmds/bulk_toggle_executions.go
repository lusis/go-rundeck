package cmds

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/lusis/go-rundeck/pkg/cli"
	rundeck "github.com/lusis/go-rundeck/pkg/rundeck"
	"github.com/spf13/cobra"
)

var (
	bulkToggleExecutionIDs = []string{}
)

func bulkToggleExecutionsFunc(cmd *cobra.Command, args []string) error {
	res := rundeck.BulkToggleResponse{}
	switch args[0] {
	case enable:
		data, err := cli.Client.BulkEnableExecution(bulkToggleExecutionIDs...)
		if err != nil {
			return err
		}
		res = *data
	case disable:
		data, err := cli.Client.BulkDisableExecution(bulkToggleExecutionIDs...)
		if err != nil {
			return err
		}
		res = *data
	}

	cli.OutputFormatter.SetHeaders([]string{
		"Request Count",
		"Enabled?",
		"All Successful",
		"Succeeded",
		"Failed",
	})
	var failureResults []string
	var successResults []string
	if len(res.Failed) != 0 {
		for _, f := range res.Failed {
			failureResults = append(failureResults, fmt.Sprintf("%s: %s", f.ID, f.Message))
		}
	}
	if len(res.Succeeded) != 0 {
		for _, f := range res.Succeeded {
			successResults = append(successResults, fmt.Sprintf("%s: %s", f.ID, f.Message))
		}
	}
	if rowErr := cli.OutputFormatter.AddRow([]string{
		strconv.Itoa(res.RequestCount),
		fmt.Sprintf("%t", res.Enabled),
		fmt.Sprintf("%t", res.AllSuccessful),
		strings.Join(successResults, "\n"),
		strings.Join(failureResults, "\n"),
	}); rowErr != nil {
		return rowErr
	}
	cli.OutputFormatter.Draw()
	return nil

}

func bulkToggleExecutionsCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "execution [enable|disable] -i [job-id][,job-id]",
		Short: "Bulk toggles executions of jobs on a rundeck server",
		RunE:  bulkToggleExecutionsFunc,
		Args:  cobra.MinimumNArgs(1),
	}
	rootCmd := cli.New(cmd)

	rootCmd.Flags().StringSliceVarP(&bulkToggleExecutionIDs, "ids", "i", nil, "list of job ids")
	_ = rootCmd.MarkFlagRequired("ids")
	return rootCmd
}
