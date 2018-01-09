package cmds

import (
	"fmt"
	"strings"

	"github.com/lusis/go-rundeck/pkg/cli"
	"github.com/spf13/cobra"
)

var (
	getProjectACLPolicyPolicyName string
)

func getProjectACLPolicyFunc(cmd *cobra.Command, args []string) error {
	projectName := args[0]
	if strings.HasSuffix(getProjectACLPolicyPolicyName, ".aclpolicy") {
		return fmt.Errorf("policy name should not end with .aclpolicy")
	}
	res, err := cli.Client.GetProjectACLPolicy(projectName, getProjectACLPolicyPolicyName)
	if err != nil {
		return err
	}
	fmt.Println(string(res))
	return nil
}

func getProjectACLPolicyCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "get project-name -p policy-name",
		Short: "gets a project acl policy",
		Args:  cobra.MinimumNArgs(1),
		RunE:  getProjectACLPolicyFunc,
	}
	rootCmd := cli.New(cmd)
	rootCmd.ResetFlags()
	rootCmd.Flags().StringVarP(&getProjectACLPolicyPolicyName, "policy-name", "p", "", "policy name to get")
	_ = rootCmd.MarkFlagRequired("policy-name")
	return rootCmd
}
