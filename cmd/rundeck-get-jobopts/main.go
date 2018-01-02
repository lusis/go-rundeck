package main

import (
	"errors"
	"fmt"

	cli "github.com/lusis/go-rundeck/pkg/cli"
	"github.com/spf13/cobra"
)

var (
	jobid string
)

func runFunc(cmd *cobra.Command, args []string) error {
	if jobid == "" {
		return errors.New("you must specify a job id")
	}
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
		Use:   "rundeck-get-jobopts -j [job-id]",
		Short: "gets an execution from the rundeck server",
		RunE:  runFunc,
	}
	cmd.Flags().StringVarP(&jobid, "job-id", "j", "", "job id")
	rootCmd := cli.New(cmd)
	_ = rootCmd.Execute()

}
