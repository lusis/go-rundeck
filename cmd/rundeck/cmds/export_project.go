package cmds

import (
	"os"
	"path/filepath"

	"github.com/lusis/go-rundeck/pkg/cli"
	rundeck "github.com/lusis/go-rundeck/pkg/rundeck"
	"github.com/spf13/cobra"
)

var (
	projectExportAll          bool
	projectExportExecutionIDs []string
	projectExportJobs         bool
	projectExportExecutions   bool
	projectExportConfigs      bool
	projectExportAcls         bool
	projectExportReadmes      bool
	projectExportFile         string
)

func exportProjectFunc(cmd *cobra.Command, args []string) error {
	projectName := args[0]
	destFile := filepath.Join(".", projectExportFile)
	f, fErr := os.Create(destFile)
	if fErr != nil {
		return fErr
	}
	opts := []rundeck.ProjectExportOption{
		rundeck.ProjectExportAll(projectExportAll),
		rundeck.ProjectExportConfigs(projectExportConfigs),
		rundeck.ProjectExportAcls(projectExportAcls),
		rundeck.ProjectExportJobs(projectExportJobs),
		rundeck.ProjectExportReadmes(projectExportReadmes),
	}
	if len(projectExportExecutionIDs) > 0 {
		opts = append(opts, rundeck.ProjectExportExecutionIDs(projectExportExecutionIDs...)) // nolint: ineffassign
	}

	return cli.Client.GetProjectArchiveExport(projectName, f, opts...)
}

func exportProjectCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "export project-name [-o destination-file]",
		Short: "exports a project from a rundeck server",
		Args:  cobra.MinimumNArgs(1),
		RunE:  exportProjectFunc,
	}
	rootCmd := cli.New(cmd)
	rootCmd.ResetFlags()

	rootCmd.Flags().StringVarP(&projectExportFile, "output-file", "o", "project-export.zip", "destination filename for the export")
	rootCmd.Flags().BoolVar(&projectExportAll, "all", true, "export everything")
	rootCmd.Flags().StringSliceVarP(&projectExportExecutionIDs, "ids", "i", nil, "specify specific execution ids to export")
	rootCmd.Flags().BoolVar(&projectExportAcls, "acls", true, "export acls")
	rootCmd.Flags().BoolVar(&projectExportConfigs, "configs", true, "export configs")
	rootCmd.Flags().BoolVar(&projectExportExecutions, "executions", true, "export executions")
	rootCmd.Flags().BoolVar(&projectExportJobs, "jobs", true, "export jobs")
	rootCmd.Flags().BoolVar(&projectExportReadmes, "readmes", true, "export readmes")
	return rootCmd
}
