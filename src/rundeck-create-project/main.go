package main

import (
	"fmt"
	"os"

	kingpin "gopkg.in/alecthomas/kingpin.v2"
	rundeck "rundeck.v17"
)

var (
	project = kingpin.Flag("project", "New project name").Required().String()
)

func main() {
	kingpin.Parse()
	client := rundeck.NewClientFromEnv()

	newProject := rundeck.NewProject{
		Name: *project,
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
