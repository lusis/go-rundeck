package cmds

import (
	cli "github.com/lusis/go-rundeck/pkg/cli"
	"github.com/spf13/cobra"
)

func deleteProjectFunc(cmd *cobra.Command, args []string) error {
	id := args[0]
	return cli.Client.DeleteProject(id)
}

func deleteProjectCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "delete project-name",
		Short: "deletes a project on the rundeck server",
		Args:  cobra.MinimumNArgs(1),
		RunE:  deleteProjectFunc,
	}
	rootCmd := cli.New(cmd)
	rootCmd.ResetFlags()
	return rootCmd
}
