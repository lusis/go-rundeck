package cmds

import (
	"os"
	"path/filepath"

	"github.com/lusis/go-rundeck/pkg/cli"
	rundeck "github.com/lusis/go-rundeck/pkg/rundeck"
	"github.com/spf13/cobra"
)

var (
	projectExportFile string
)

func exportProjectFunc(cmd *cobra.Command, args []string) error {
	projectName := args[0]
	destFile := filepath.Join(".", projectExportFile)
	f, fErr := os.Create(destFile)
	if fErr != nil {
		return fErr
	}
	return cli.Client.GetProjectArchiveExport(projectName, f, rundeck.ProjectExportAll(true))
}

func exportProjectCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "export project-name [-f destination-file]",
		Short: "exports a project from a rundeck server",
		Args:  cobra.MinimumNArgs(1),
		RunE:  exportProjectFunc,
	}
	rootCmd := cli.New(cmd)
	rootCmd.ResetFlags()
	rootCmd.Flags().StringVarP(&projectExportFile, "file", "f", "project-export.zip", "destination filename for the export")
	return rootCmd
}
