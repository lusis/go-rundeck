package cmds

import (
	"github.com/lusis/go-rundeck/pkg/cli"
	"github.com/spf13/cobra"
)

func setupProjectSCMFunc(cmd *cobra.Command, args []string) error {
	projectName := args[0]
	integration := args[1]
	pluginType := args[2]

	params, paramsErr := ParseSliceKeyValue(scmSetupParams)
	if paramsErr != nil {
		return paramsErr
	}
	_, err := cli.Client.SetupSCMPluginForProject(projectName, integration, pluginType, params)
	return err
}
func setupProjectSCMCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "setup project-name integration plugin-type -o setting=value [-o setting=value]",
		Short:   "sets up scm integration for a rundeck project",
		Example: "setup fooprj import git-import -o dir=/var/tmp/fooprj -o url=/home/rundeck-import.git/ -o pathTemplate=\\${job.group}\\${job.name}-\\${job.id}.\\${config.format} -o branch=master -o format=yaml -o strictHostKeyChecking=no",
		Args:    cobra.MinimumNArgs(3),
		RunE:    setupProjectSCMFunc,
	}

	rootCmd := cli.New(cmd)
	rootCmd.Flags().StringSliceVarP(&scmSetupParams, "option", "o", []string{}, "repeatable list of key/value options in the format of key=value")
	return rootCmd
}
