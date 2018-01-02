package main

import (
	"fmt"

	"github.com/lusis/go-rundeck/pkg/cli"
	"github.com/spf13/cobra"
)

var (
	jobformat string
)

func runFunc(cmd *cobra.Command, args []string) error {
	jobid := args[0]

	res, err := cli.Client.ExportJob(jobid, jobformat)
	if err != nil {
		return err
	}
	fmt.Println(string(res))
	return nil
}
func main() {
	cmd := &cobra.Command{
		Use:   "rundeck-export-job jobid [-f format]",
		Short: "exports a job in the specified format",
		Args:  cobra.MinimumNArgs(1),
		RunE:  runFunc,
	}
	cmd.Flags().StringVarP(&jobformat, "job-format", "f", "yaml", "format to export job")
	rootCmd := cli.New(cmd)
	_ = rootCmd.Execute()
}
