package cmds

import (
	"github.com/spf13/cobra"
)

func projectPoliciesCommands() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "policy",
		Short: "operate on rundeck project acl policies",
	}
	cmd.AddCommand(getProjectACLPolicyCommand())
	cmd.AddCommand(deleteProjectACLPolicyCommand())
	cmd.AddCommand(updateProjectACLPolicyCommand())
	cmd.AddCommand(createProjectACLPolicyCommand())
	cmd.AddCommand(deleteProjectACLPolicyCommand())
	cmd.AddCommand(listProjectACLPoliciesCommand())
	return cmd
}

func projectCommands() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "project",
		Short: "operate on a rundeck project",
	}
	cmd.AddCommand(getProjectCommand())
	cmd.AddCommand(deleteProjectCommand())
	cmd.AddCommand(createProjectCommand())
	cmd.AddCommand(getJobsCommand())
	cmd.AddCommand(projectExecutionsCommand())
	cmd.AddCommand(projectHistoryCommand())
	cmd.AddCommand(getProjectConfigCommand())
	cmd.AddCommand(exportProjectCommand())
	cmd.AddCommand(projectPoliciesCommands())
	cmd.AddCommand(scmCommands())
	return cmd
}
