package main

import (
	"errors"
	"fmt"

	cli "github.com/lusis/go-rundeck/pkg/cli"
	"github.com/spf13/cobra"
)

var (
	id string
)

func runFunc(cmd *cobra.Command, args []string) error {
	if id == "" {
		return errors.New("you must specify an execution id")
	}
	data, err := cli.Client.GetExecution(id)
	if err != nil {
		return err
	}
	cli.OutputFormatter.SetHeaders([]string{
		"ID",
		"User",
		"Status",
		"Start Date",
		"End Date",
		"Args",
		"Server UUID",
		"Custom Status",
	})
	if rowErr := cli.OutputFormatter.AddRow([]string{
		fmt.Sprintf("%d", data.ID),
		data.User,
		data.Status,
		data.DateStarted.Date.String(),
		data.DateEnded.Date.String(),
		data.ArgString,
		data.ServerUUID,
		data.CustomStatus,
	}); rowErr != nil {
		return rowErr
	}

	cli.OutputFormatter.Draw()
	return nil
}
func main() {
	cmd := &cobra.Command{
		Use:   "rundeck-get-execution -e [execution-id]",
		Short: "gets an execution from the rundeck server",
		RunE:  runFunc,
	}
	cmd.Flags().StringVarP(&id, "id", "i", "", "execution id")
	rootCmd := cli.New(cmd)
	_ = rootCmd.Execute()
}
