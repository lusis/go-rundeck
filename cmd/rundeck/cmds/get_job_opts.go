package cmds

import (
	"fmt"

	"github.com/lusis/go-rundeck/pkg/cli"
	"github.com/spf13/cobra"
)

func getJobOptsFunc(cmd *cobra.Command, args []string) error {
	jobid := args[0]
	data, err := cli.Client.GetJobOpts(jobid)
	if err != nil {
		return err
	}
	cli.OutputFormatter.SetHeaders([]string{
		"Name",
		"Description",
		"Value",
		"Required?",
		"Regex",
	})
	for _, d := range data {
		if err := cli.OutputFormatter.AddRow([]string{
			d.Name,
			d.Description,
			d.Value,
			fmt.Sprintf("%t", d.Required),
			d.Regex}); err != nil {
			return err
		}
	}
	cli.OutputFormatter.Draw()
	return nil
}

func getJobOptsCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "opts job-id",
		Short: "gets a job's options from a rundeck server",
		Args:  cobra.MinimumNArgs(1),
		RunE:  getJobOptsFunc,
	}
	rootCmd := cli.New(cmd)
	return rootCmd
}
