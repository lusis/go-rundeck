package cmds

import (
	"fmt"
	"strings"

	"github.com/lusis/go-rundeck/pkg/cli"
	"github.com/spf13/cobra"
)

var (
	deleteProjectACLPolicyPolicyName string
)

func deleteProjectACLPolicyFunc(cmd *cobra.Command, args []string) error {
	projectName := args[0]
	if strings.HasSuffix(deleteProjectACLPolicyPolicyName, ".aclpolicy") {
		return fmt.Errorf("policy name should not end with .aclpolicy")
	}
	return cli.Client.DeleteProjectACLPolicy(projectName, deleteProjectACLPolicyPolicyName)
}

func deleteProjectACLPolicyCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "delete project-name -p policy-name",
		Short: "deletes a project acl policy",
		Args:  cobra.MinimumNArgs(1),
		RunE:  deleteProjectACLPolicyFunc,
	}
	rootCmd := cli.New(cmd)
	rootCmd.ResetFlags()
	rootCmd.Flags().StringVarP(&deleteProjectACLPolicyPolicyName, "policy-name", "p", "", "policy name to get")
	_ = rootCmd.MarkFlagRequired("policy-name")
	return rootCmd
}
