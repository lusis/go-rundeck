package cmds

import "github.com/spf13/cobra"

func policiesCommands() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "policy",
		Short: "operate on rundeck acl policies",
	}
	cmd.AddCommand(createACLPolicyCommand())
	return cmd
}
