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
	importProjectACLPolicyFile     string
	importProjectACLPolicyFileName string
)

func createProjectACLPolicyFunc(cmd *cobra.Command, args []string) error {
	projectName := args[0]
	if strings.HasSuffix(importProjectACLPolicyFile, ".aclpolicy") {
		return fmt.Errorf("policy name should not end with .aclpolicy")
	}
	policyFile, err := os.Open(importProjectACLPolicyFileName)
	if err != nil {
		return err
	}
	contents, readErr := ioutil.ReadAll(policyFile)
	if readErr != nil {
		return readErr
	}
	if err := cli.Client.CreateProjectACLPolicy(projectName, importProjectACLPolicyFile, bytes.NewReader(contents)); err != nil {
		return err
	}
	fmt.Println("policy created")
	return nil
}

func createProjectACLPolicyCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "create policy-name -p policy-name -f policy-file",
		Short: "creates a project acl policy on the rundeck server. must be a valid yaml ACL policy",
		Args:  cobra.MinimumNArgs(1),
		RunE:  createProjectACLPolicyFunc,
	}
	rootCmd := cli.New(cmd)
	rootCmd.ResetFlags()
	rootCmd.Flags().StringVarP(&importProjectACLPolicyFile, "policy-name", "p", "", "name of the policy to create")
	rootCmd.Flags().StringVarP(&importProjectACLPolicyFileName, "file", "f", "", "full path to policy file")
	_ = rootCmd.MarkFlagRequired("file")
	_ = rootCmd.MarkFlagRequired("policy-name")

	return rootCmd
}

func updateProjectACLPolicyFunc(cmd *cobra.Command, args []string) error {
	projectName := args[0]
	if strings.HasSuffix(importProjectACLPolicyFile, ".aclpolicy") {
		return fmt.Errorf("policy name should not end with .aclpolicy")
	}
	policyFile, err := os.Open(importProjectACLPolicyFileName)
	if err != nil {
		return err
	}
	contents, readErr := ioutil.ReadAll(policyFile)
	if readErr != nil {
		return readErr
	}
	if err := cli.Client.UpdateProjectACLPolicy(projectName, importProjectACLPolicyFile, bytes.NewReader(contents)); err != nil {
		return err
	}
	fmt.Println("policy updated")
	return nil
}

func updateProjectACLPolicyCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "update project-name -p policy-name -f policy-file",
		Short: "updates a project acl policy on the rundeck server. must be a valid yaml ACL policy",
		Args:  cobra.MinimumNArgs(1),
		RunE:  updateProjectACLPolicyFunc,
	}
	rootCmd := cli.New(cmd)
	rootCmd.ResetFlags()
	rootCmd.Flags().StringVarP(&importProjectACLPolicyFile, "policy-name", "p", "", "name of the policy to create")
	rootCmd.Flags().StringVarP(&importProjectACLPolicyFileName, "file", "f", "", "full path to policy file")
	_ = rootCmd.MarkFlagRequired("file")
	_ = rootCmd.MarkFlagRequired("policy-name")

	return rootCmd
}
