package cmds

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"os"
	"strings"

	cli "github.com/lusis/go-rundeck/pkg/cli"
	"github.com/spf13/cobra"
)

var (
	importSystemACLPolicyFileName string
)

func createSystemACLPolicyFunc(cmd *cobra.Command, args []string) error {
	policyName := args[0]
	if strings.HasSuffix(policyName, ".aclpolicy") {
		return fmt.Errorf("policy name should not end with .aclpolicy")
	}
	policyFile, err := os.Open(importSystemACLPolicyFileName)
	if err != nil {
		return err
	}
	contents, readErr := ioutil.ReadAll(policyFile)
	if readErr != nil {
		return readErr
	}
	if err := cli.Client.CreateSystemACLPolicy(policyName, bytes.NewReader(contents)); err != nil {
		return err
	}
	fmt.Println("policy created")
	return nil
}

func createSystemACLPolicyCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "create policy-name -f policy-file",
		Short: "creates an acl policy on the rundeck server. must be a valid yaml ACL policy",
		Args:  cobra.MinimumNArgs(1),
		RunE:  createSystemACLPolicyFunc,
	}
	rootCmd := cli.New(cmd)
	rootCmd.ResetFlags()
	rootCmd.Flags().StringVarP(&importSystemACLPolicyFileName, "file", "f", "", "full path to policy file")
	_ = rootCmd.MarkFlagRequired("file")

	return rootCmd
}

func updateSystemACLPolicyFunc(cmd *cobra.Command, args []string) error {
	policyName := args[0]
	if strings.HasSuffix(policyName, ".aclpolicy") {
		return fmt.Errorf("policy name should not end with .aclpolicy")
	}
	policyFile, err := os.Open(importSystemACLPolicyFileName)
	if err != nil {
		return err
	}
	contents, readErr := ioutil.ReadAll(policyFile)
	if readErr != nil {
		return readErr
	}
	if err := cli.Client.UpdateSystemACLPolicy(policyName, bytes.NewReader(contents)); err != nil {
		return err
	}
	fmt.Println("policy update")
	return nil
}

func updateSystemACLPolicyCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "update policy-name -f policy-file",
		Short: "updates an acl policy on the rundeck server. must be a valid yaml ACL policy",
		Args:  cobra.MinimumNArgs(1),
		RunE:  updateSystemACLPolicyFunc,
	}
	rootCmd := cli.New(cmd)
	rootCmd.ResetFlags()
	rootCmd.Flags().StringVarP(&importSystemACLPolicyFileName, "file", "f", "", "full path to policy file")
	_ = rootCmd.MarkFlagRequired("file")

	return rootCmd
}
