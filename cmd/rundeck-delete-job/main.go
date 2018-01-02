package main

import (
	"errors"

	cli "github.com/lusis/go-rundeck/pkg/cli"
	"github.com/spf13/cobra"
)

var id string

func runFunc(cmd *cobra.Command, args []string) error {
	if id == "" {
		return errors.New("an id is required")
	}
	return cli.Client.DeleteJob(id)
}

func main() {
	cmd := &cobra.Command{
		Use:   "rundeck-delete-job -j id",
		Short: "deletes a job on the rundeck server",
		RunE:  runFunc,
	}
	cmd.Flags().StringVarP(&id, "job-id", "j", "", "job id")
	cli.UseFormatter = false
	rootCmd := cli.New(cmd)
	_ = rootCmd.Execute()
}
