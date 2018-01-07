package cmds

import (
	"github.com/lusis/go-rundeck/pkg/cli"
	"github.com/spf13/cobra"
)

func deleteSystemACLPolicyFunc(cmd *cobra.Command, args []string) error {
	policyName := args[0]
	return cli.Client.DeleteSystemACLPolicy(policyName)
}

func deleteSystemACLPolicyCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "delete policy-name",
		Short: "deletes a system acl policy",
		Args:  cobra.MinimumNArgs(1),
		RunE:  deleteSystemACLPolicyFunc,
	}
	rootCmd := cli.New(cmd)
	rootCmd.ResetFlags()
	return rootCmd
}
