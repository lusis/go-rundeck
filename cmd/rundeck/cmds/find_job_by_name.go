package cmds

import (
	"fmt"

	"github.com/lusis/go-rundeck/pkg/cli"
	"github.com/spf13/cobra"
)

func findJobByNameFunc(cmd *cobra.Command, args []string) error {
	jobname := args[0]

	res, err := cli.Client.FindJobByName(jobname)
	if err != nil {
		return err
	}
	cli.OutputFormatter.SetHeaders([]string{
		"ID",
		"Name",
		"Project",
		"Description",
		"Group",
		"Scheduled?",
		"Schedule Enabled?",
		"Enabled?",
		"Average Duration",
	})
	for _, data := range res {
		if rowErr := cli.OutputFormatter.AddRow([]string{
			data.ID,
			data.Name,
			data.Project,
			data.Description,
			data.Group,
			fmt.Sprintf("%t", data.Scheduled),
			fmt.Sprintf("%t", data.ScheduleEnabled),
			fmt.Sprintf("%t", data.Enabled),
			fmt.Sprintf("%d", data.AverageDuration),
		}); rowErr != nil {
			return rowErr
		}
	}
	cli.OutputFormatter.Draw()
	return nil
}

func findJobByNameCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "find job-name",
		Short: "finds a project's job by name",
		Args:  cobra.MinimumNArgs(1),
		RunE:  findJobByNameFunc,
	}
	rootCmd := cli.New(cmd)
	return rootCmd
}
