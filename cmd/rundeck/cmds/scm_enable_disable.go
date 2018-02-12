package cmds

import (
	"github.com/lusis/go-rundeck/pkg/cli"
	"github.com/spf13/cobra"
)

func enableProjectSCMFunc(cmd *cobra.Command, args []string) error {
	projectName := args[0]
	integration := args[1]
	pluginType := args[2]
	return cli.Client.EnableSCMPluginForProject(projectName, integration, pluginType)
}

func disableProjectSCMFunc(cmd *cobra.Command, args []string) error {
	projectName := args[0]
	integration := args[1]
	pluginType := args[2]
	return cli.Client.DisableSCMPluginForProject(projectName, integration, pluginType)
}

func enableProjectSCMCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "enable integration plugin-type",
		Short: "enables a project's scm configuration",
		Args:  cobra.MinimumNArgs(3),
		RunE:  enableProjectSCMFunc,
	}
	rootCmd := cli.New(cmd)
	rootCmd.ResetFlags()
	return rootCmd
}

func disableProjectSCMCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "disable project-name integration plugin-type",
		Short: "disables a project's scm configuration",
		Args:  cobra.MinimumNArgs(3),
		RunE:  disableProjectSCMFunc,
	}
	rootCmd := cli.New(cmd)
	rootCmd.ResetFlags()
	return rootCmd
}
