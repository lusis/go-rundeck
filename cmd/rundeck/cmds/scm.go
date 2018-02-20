package cmds

import (
	"github.com/spf13/cobra"
)

var scmInputFieldsFullDescription bool
var scmInputFieldsRequiredOnly bool
var scmSetupParams []string

func scmCommands() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "scm",
		Short: "operate on rundeck scm plugins",
	}
	cmd.AddCommand(getSCMPluginInputFieldsCommand())
	cmd.AddCommand(enableProjectSCMCommand())
	cmd.AddCommand(disableProjectSCMCommand())
	cmd.AddCommand(listProjectSCMPluginsCommand())
	cmd.AddCommand(setupProjectSCMCommand())
	return cmd
}
