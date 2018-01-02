package main

import (
	cli "github.com/lusis/go-rundeck/pkg/cli"
	"github.com/spf13/cobra"
)

func runFunc(cmd *cobra.Command, args []string) error {
	data, err := cli.Client.ListProjects()
	if err != nil {
		return err
	}

	cli.OutputFormatter.SetHeaders([]string{
		"Name",
		"Description",
		"URL",
	})
	for _, d := range *data {
		if rowErr := cli.OutputFormatter.AddRow([]string{
			d.Name,
			d.Description,
			d.URL,
		}); rowErr != nil {
			return rowErr
		}
	}
	cli.OutputFormatter.Draw()
	return nil
}
func main() {
	cmd := &cobra.Command{
		Use:   "rundeck-list-projects",
		Short: "gets a list of projects from the rundeck server",
		RunE:  runFunc,
	}
	rootCmd := cli.New(cmd)
	_ = rootCmd.Execute()
}
