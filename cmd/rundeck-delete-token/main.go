package main

import (
	cli "github.com/lusis/go-rundeck/pkg/cli"
	"github.com/spf13/cobra"
)

func runFunc(cmd *cobra.Command, args []string) error {
	id := args[0]
	return cli.Client.DeleteToken(id)
}

func main() {
	cmd := &cobra.Command{
		Use:   "rundeck-delete-Token token-id",
		Short: "deletes an api on the rundeck server",
		Args:  cobra.MinimumNArgs(1),
		RunE:  runFunc,
	}
	cli.UseFormatter = false
	rootCmd := cli.New(cmd)
	_ = rootCmd.Execute()
}
