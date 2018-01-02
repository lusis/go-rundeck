package main

import (
	"errors"

	cli "github.com/lusis/go-rundeck/pkg/cli"
	"github.com/spf13/cobra"
)

var (
	projectid string
)

func runFunc(cmd *cobra.Command, args []string) error {
	if projectid == "" {
		return errors.New("you must specify a project id")
	}
	data, err := cli.Client.ListJobs(projectid)
	if err != nil {
		return err
	}
	cli.OutputFormatter.SetHeaders([]string{"ID", "Name", "Description", "Group", "Project"})
	for _, d := range *data {
		if err := cli.OutputFormatter.AddRow([]string{d.ID, d.Name, d.Description, d.Group, d.Project}); err != nil {
			return err
		}
	}
	cli.OutputFormatter.Draw()
	return nil

}
func main() {
	cmd := &cobra.Command{
		Use:   "rundeck-list-jobs -p project-id",
		Short: "gets a list of jobs for a project from the rundeck server",
		RunE:  runFunc,
	}
	cmd.Flags().StringVarP(&projectid, "project-id", "p", "", "project id")
	rootCmd := cli.New(cmd)
	_ = rootCmd.Execute()
}
