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
	return cli.Client.DeleteExecution(id)
}

func main() {
	cmd := &cobra.Command{
		Use:   "rundeck-delete-execution -e id",
		Short: "deletes an execution on the rundeck server",
		RunE:  runFunc,
	}
	cmd.Flags().StringVarP(&id, "execution-id", "e", "", "execution id")
	rootCmd := cli.New(cmd)
	_ = rootCmd.Execute()
}
