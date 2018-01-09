package cmds

import (
	"fmt"
	"strings"

	"github.com/lusis/go-rundeck/pkg/cli"
	"github.com/spf13/cobra"
)

func getSystemACLPolicyFunc(cmd *cobra.Command, args []string) error {
	policyName := args[0]
	if strings.HasSuffix(policyName, ".aclpolicy") {
		return fmt.Errorf("policy name should not end with .aclpolicy")
	}
	res, err := cli.Client.GetSystemACLPolicy(policyName)
	if err != nil {
		return err
	}
	fmt.Println(string(res))
	return nil
}

func getSystemACLPolicyCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "get policy-name",
		Short: "gets a system acl policy",
		Args:  cobra.MinimumNArgs(1),
		RunE:  getSystemACLPolicyFunc,
	}
	rootCmd := cli.New(cmd)
	rootCmd.ResetFlags()
	return rootCmd
}
