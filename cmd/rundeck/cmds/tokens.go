package cmds

import "github.com/spf13/cobra"

func tokensCommands() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "tokens",
		Short: "operate on rundeck api tokens",
	}
	cmd.AddCommand(getTokensCommand())
	return cmd
}
