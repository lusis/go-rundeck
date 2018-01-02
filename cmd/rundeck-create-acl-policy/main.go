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
	filename string
)

func runFunc(cmd *cobra.Command, args []string) error {
	policyName := args[0]
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
		Use:   "rundeck-create-acl-policy policy-name -f policy-file",
		Short: "creates an acl policy on the rundeck server. must be a valid yaml ACL policy",
		Args:  cobra.MinimumNArgs(1),
		RunE:  runFunc,
	}
	cmd.Flags().StringVarP(&filename, "file", "f", "", "full path to policy file")
	_ = cmd.MarkFlagRequired("file")
	cli.UseFormatter = false
	rootCmd := cli.New(cmd)
	_ = rootCmd.Execute()
}
