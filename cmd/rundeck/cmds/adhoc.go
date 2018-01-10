package cmds

import "github.com/spf13/cobra"

var (
	adHocAsUser            string
	adHocFilter            string
	adHocNodeThreadCount   int
	adHocNodeKeepGoing     bool
	adHocScriptInterpreter string
	adHocArgString         string
	adHocArgsQuoted        bool
	adHocFileExtension     string
)

func adHocCommands() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "adhoc",
		Short: "run adhoc commands, scripts and scripts from urls against a project",
	}
	cmd.PersistentFlags().StringVar(&adHocAsUser, "as-user", "", "rundeck user to run as")
	cmd.PersistentFlags().StringVar(&adHocFilter, "filter", "", "rundeck filter to use")
	cmd.PersistentFlags().IntVar(&adHocNodeThreadCount, "thread-count", 1, "node thread count")
	cmd.PersistentFlags().BoolVar(&adHocNodeKeepGoing, "keep-going", false, "keep going on failure")
	cmd.AddCommand(runAdHocCmdCommand())
	cmd.AddCommand(runAdHocScriptCommand())
	cmd.AddCommand(runAdHocURLCommand())
	return cmd
}
