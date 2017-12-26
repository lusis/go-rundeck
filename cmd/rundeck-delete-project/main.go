package main

import (
	"fmt"
	"log"
	"os"

	rundeck "github.com/lusis/go-rundeck/pkg/rundeck.v21"
	kingpin "gopkg.in/alecthomas/kingpin.v2"
)

var (
	project = kingpin.Arg("project", "Project to delete").Required().String()
)

func main() {
	kingpin.Parse()
	client, clientErr := rundeck.NewClientFromEnv()
	if clientErr != nil {
		log.Fatal(clientErr.Error())
	}

	err := client.DeleteProject(*project)
	if err != nil {
		fmt.Printf("%s\n", err)
		os.Exit(1)
	} else {
		fmt.Printf("Project %s deleted\n", *project)
		os.Exit(0)
	}
}
