package cmds

import (
	"github.com/spf13/cobra"
)

func tokenCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "token",
		Short: "operate on an individual token in rundeck",
	}
	cmd.AddCommand(createTokenCommand())
	cmd.AddCommand(deleteTokenCommand())
	return cmd
}
