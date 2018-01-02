package main

import (
	"fmt"

	cli "github.com/lusis/go-rundeck/pkg/cli"
	"github.com/spf13/cobra"
)

func runFunc(cmd *cobra.Command, args []string) error {
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
func main() {
	cmd := &cobra.Command{
		Use:   "rundeck-get-jobopts job-id",
		Short: "gets an execution from the rundeck server",
		Args:  cobra.MinimumNArgs(1),
		RunE:  runFunc,
	}
	rootCmd := cli.New(cmd)
	_ = rootCmd.Execute()

}
