package cmds

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"os"

	cli "github.com/lusis/go-rundeck/pkg/cli"
	"github.com/spf13/cobra"
)

var (
	filename string
)

func createACLPolicyFunc(cmd *cobra.Command, args []string) error {
	policyName := args[0]
	policyFile, err := os.Open(filename)
	if err != nil {
		return err
	}
	contents, readErr := ioutil.ReadAll(policyFile)
	if readErr != nil {
		return readErr
	}
	if err := cli.Client.CreateACLPolicy(policyName, bytes.NewReader(contents)); err != nil {
		return err
	}
	fmt.Println("policy created")
	return nil
}

func createACLPolicyCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "create policy-name -f policy-file",
		Short: "creates an acl policy on the rundeck server. must be a valid yaml ACL policy",
		Args:  cobra.MinimumNArgs(1),
		RunE:  createACLPolicyFunc,
	}
	rootCmd := cli.New(cmd)
	rootCmd.ResetFlags()
	rootCmd.Flags().StringVarP(&filename, "file", "f", "", "full path to policy file")
	_ = rootCmd.MarkFlagRequired("file")

	return rootCmd
}
