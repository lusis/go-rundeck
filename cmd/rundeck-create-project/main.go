package main

import (
	"fmt"
	"os"

	rundeck "github.com/lusis/go-rundeck/pkg/rundeck.v19"
	kingpin "gopkg.in/alecthomas/kingpin.v2"
)

var (
	project     = kingpin.Arg("project", "New project name").Required().String()
	description = kingpin.Flag("description", "Description of the project").String()
)

func main() {
	kingpin.Parse()
	client := rundeck.NewClientFromEnv()

	newProject := rundeck.NewProject{
		Name: *project,
	}
	if description != nil {
		newProject.Description = *description
	}
	err := client.MakeProject(newProject)
	if err != nil {
		fmt.Printf("%s\n", err)
		os.Exit(1)
	} else {
		fmt.Printf("Project %s created\n", *project)
		os.Exit(0)
	}
}
