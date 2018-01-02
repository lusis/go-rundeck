package main

import (
	"bytes"
	"errors"
	"io/ioutil"
	"os"

	cli "github.com/lusis/go-rundeck/pkg/cli"
	"github.com/spf13/cobra"
)

var (
	policyName string
	filename   string
)

func runFunc(cmd *cobra.Command, args []string) error {
	if policyName == "" || filename == "" {
		return errors.New("you must specify both a policy name and a filename")
	}
	policyFile, err := os.Open(filename)
	if err != nil {
		return err
	}
	contents, _ := ioutil.ReadAll(policyFile)
	return cli.Client.CreateACLPolicy(policyName, bytes.NewReader(contents))
}

func main() {
	cmd := &cobra.Command{
		Use:   "rundeck-create-acl-policy -f policy-file -n policy-name",
		Short: "creates an acl policy on the rundeck server. must be a valid yaml ACL policy",
		RunE:  runFunc,
	}
	cmd.Flags().StringVarP(&policyName, "policy-name", "n", "", "policy name")
	cmd.Flags().StringVarP(&filename, "file", "f", "", "full path to policy file")
	rootCmd := cli.New(cmd)
	_ = rootCmd.Execute()
}
