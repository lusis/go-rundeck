package cmds

import (
	"fmt"
	"os"

	"github.com/lusis/go-rundeck/pkg/cli"
	rundeck "github.com/lusis/go-rundeck/pkg/rundeck"
	"github.com/spf13/cobra"
)

var (
	importJobFormat     string
	importJobDupeOption string
	importJobUUIDOption string
	importJobProject    string
)

func importJobFunc(cmd *cobra.Command, args []string) error {
	jobfile := args[0]
	file, fileErr := os.Open(jobfile)
	defer file.Close() // nolint: errcheck
	if fileErr != nil {
		return fileErr
	}
	res, err := cli.Client.ImportJob(importJobProject, file,
		rundeck.ImportFormat(importJobFormat),
		rundeck.ImportDupe(importJobDupeOption),
		rundeck.ImportUUID(importJobUUIDOption))
	if err != nil {
		return err
	}
	cli.OutputFormatter.SetHeaders([]string{
		"Result",
		"Index",
		"ID",
		"Name",
		"Group",
		"Project",
		"HRef",
		"Permalink",
		"Messages",
	})
	for _, r := range res.Succeeded {
		if rowErr := cli.OutputFormatter.AddRow([]string{
			"Succeeded",
			fmt.Sprintf("%d", r.Index),
			r.ID,
			r.Name,
			r.Group,
			r.Project,
			r.HRef,
			r.Permalink,
			r.Messages,
		}); rowErr != nil {
			return rowErr
		}
	}
	for _, r := range res.Failed {
		if rowErr := cli.OutputFormatter.AddRow([]string{
			"Failed",
			fmt.Sprintf("%d", r.Index),
			r.ID,
			r.Name,
			r.Group,
			r.Project,
			r.HRef,
			r.Permalink,
			r.Messages,
		}); rowErr != nil {
			return rowErr
		}
	}
	for _, r := range res.Skipped {
		if rowErr := cli.OutputFormatter.AddRow([]string{
			"Skipped",
			fmt.Sprintf("%d", r.Index),
			r.ID,
			r.Name,
			r.Group,
			r.Project,
			r.HRef,
			r.Permalink,
			r.Messages,
		}); rowErr != nil {
			return rowErr
		}
	}
	cli.OutputFormatter.Draw()
	return nil
}

func importJobCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "import job-definition -p project-name [-d dupe-handling] [-u uuid-handling] [-f format]",
		Short: "imports a job definition into a rundeck server",
		Args:  cobra.MinimumNArgs(1),
		RunE:  importJobFunc,
	}

	rootCmd := cli.New(cmd)
	rootCmd.Flags().StringVarP(&importJobFormat, "job-format", "f", "yaml", "format of job import file")
	rootCmd.Flags().StringVarP(&importJobDupeOption, "dupes", "d", "create", "how to handle existing jobs with the same name [create|update|skip]")
	rootCmd.Flags().StringVarP(&importJobUUIDOption, "uuids", "u", "preserve", "preserve or strip uuids")
	rootCmd.Flags().StringVarP(&importJobProject, "project", "p", "", "project to import the job into")
	_ = rootCmd.MarkFlagRequired("project")
	return rootCmd
}
