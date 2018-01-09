package cmds

import "github.com/spf13/cobra"

func systemPoliciesCommands() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "system-policy",
		Short: "operate on rundeck system acl policies",
	}
	cmd.AddCommand(getSystemACLPolicyCommand())
	cmd.AddCommand(listSystemACLPoliciesCommand())
	cmd.AddCommand(createSystemACLPolicyCommand())
	cmd.AddCommand(updateSystemACLPolicyCommand())
	cmd.AddCommand(deleteSystemACLPolicyCommand())
	return cmd
}
