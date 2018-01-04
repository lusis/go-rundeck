package main

import (
	"fmt"

	cli "github.com/lusis/go-rundeck/pkg/cli"
	"github.com/spf13/cobra"
)

func runFunc(cmd *cobra.Command, args []string) error {
	id := args[0]
	res, err := cli.Client.EnableExecution(id)
	if err != nil {
		return err
	}
	fmt.Printf("%t\n", res)
	return nil

}

func main() {
	cmd := &cobra.Command{
		Use:   "rundeck-enable-execution job-id",
		Short: "enables a job's execution on the rundeck server",
		Args:  cobra.MinimumNArgs(1),
		RunE:  runFunc,
	}
	cli.UseFormatter = false
	rootCmd := cli.New(cmd)
	_ = rootCmd.Execute()
}
