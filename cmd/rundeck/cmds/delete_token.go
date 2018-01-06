package cmds

import (
	cli "github.com/lusis/go-rundeck/pkg/cli"
	"github.com/spf13/cobra"
)

func deleteTokenFunc(cmd *cobra.Command, args []string) error {
	id := args[0]
	return cli.Client.DeleteToken(id)
}

func deleteTokenCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "delete token-id",
		Short: "deletes an api token on the rundeck server",
		Args:  cobra.MinimumNArgs(1),
		RunE:  deleteTokenFunc,
	}

	rootCmd := cli.New(cmd)
	rootCmd.ResetFlags()
	return rootCmd
}
