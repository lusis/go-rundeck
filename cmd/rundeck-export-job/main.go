package main

import (
	"errors"
	"fmt"

	"github.com/lusis/go-rundeck/pkg/cli"
	"github.com/spf13/cobra"
)

var (
	jobid     string
	jobformat string
)

func runFunc(cmd *cobra.Command, args []string) error {
	if jobid == "" {
		return errors.New("you must specify a jobid")
	}

	res, err := cli.Client.ExportJob(jobid, jobformat)
	if err != nil {
		return err
	}
	fmt.Printf("%s\n", string(res))
	return nil
}
func main() {
	cmd := &cobra.Command{
		Use:   "rundeck-export-job -j jobid [-f format]",
		Short: "exports a job in the specified format",
		RunE:  runFunc,
	}
	cmd.Flags().StringVarP(&jobid, "job-id", "j", "", "jobid to export")
	cmd.Flags().StringVarP(&jobformat, "job-format", "f", "yaml", "format to export job")
	rootCmd := cli.New(cmd)
	_ = rootCmd.Execute()
}
