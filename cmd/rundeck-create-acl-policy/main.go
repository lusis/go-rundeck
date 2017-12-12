package main

import (
	"fmt"
	"io/ioutil"
	"os"

	rundeck "github.com/lusis/go-rundeck/pkg/rundeck.v19"
	kingpin "gopkg.in/alecthomas/kingpin.v2"
)

var (
	policyName = kingpin.Arg("policy_name", "name for the new policy. file extension of '.aclpolicy' will be appended").Required().String()
	filename   = kingpin.Flag("policy_file", "Full /path/to/policy/file to import. Must be yaml").Required().ExistingFile()
)

func main() {
	kingpin.Parse()
	client := rundeck.NewClientFromEnv()
	policyFile, err := os.Open(*filename)
	if err != nil {
		fmt.Printf("Unable to read policy file: %s\n", err.Error())
	}
	contents, _ := ioutil.ReadAll(policyFile)
	err = client.CreateSystemACLPolicy(*policyName, contents)
	if err != nil {
		fmt.Printf("%s\n", err)
		os.Exit(1)
	} else {
		fmt.Printf("Policy created\n")
		os.Exit(0)
	}
}
