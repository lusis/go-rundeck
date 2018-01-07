package cmds

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/lusis/go-rundeck/pkg/rundeck/responses"

	"github.com/lusis/go-rundeck/pkg/cli"
	"github.com/spf13/cobra"
)

var (
	bulkToggleScheduleIDs = []string{}
)

func bulkToggleScheduleFunc(cmd *cobra.Command, args []string) error {
	res := &responses.BulkToggleResponse{}
	switch args[0] {
	case enable:
		data, err := cli.Client.BulkEnableSchedule(bulkToggleScheduleIDs...)
		if err != nil {
			return err
		}
		res = data
	case disable:
		data, err := cli.Client.BulkDisableSchedule(bulkToggleScheduleIDs...)
		if err != nil {
			return err
		}
		res = data
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

func bulkToggleScheduleCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "schedule [enable|disable] -i [job-id][,job-id]",
		Short: "Bulk toggles scheduled executions of jobs on a rundeck server",
		RunE:  bulkToggleScheduleFunc,
		Args:  cobra.MinimumNArgs(1),
	}
	rootCmd := cli.New(cmd)

	rootCmd.Flags().StringSliceVarP(&bulkToggleScheduleIDs, "ids", "i", nil, "list of job ids")
	_ = rootCmd.MarkFlagRequired("ids")
	return rootCmd
}
