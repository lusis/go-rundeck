package cmds

import (
	"fmt"

	"github.com/lusis/go-rundeck/pkg/cli"
	"github.com/spf13/cobra"
)

var exportJobFormat string

func exportJobFunc(cmd *cobra.Command, args []string) error {
	jobid := args[0]

	res, err := cli.Client.ExportJob(jobid, exportJobFormat)
	if err != nil {
		return err
	}
	fmt.Println(string(res))
	return nil
}

func exportJobCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "export jobid [-f format]",
		Short: "exports a job in the specified format",
		Args:  cobra.MinimumNArgs(1),
		RunE:  exportJobFunc,
	}

	rootCmd := cli.New(cmd)
	rootCmd.ResetFlags()
	rootCmd.Flags().StringVarP(&exportJobFormat, "job-format", "f", "yaml", "format to export job")
	return rootCmd
}
