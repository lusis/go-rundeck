package cmds

import (
	"github.com/spf13/cobra"
)

func httpCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "http",
		Short: "perform authenticated http operations against a rundeck server. kinda like curl",
	}
	cmd.AddCommand(httpGetCommand())
	return cmd
}
