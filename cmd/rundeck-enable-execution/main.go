package main

import (
	"fmt"
	"log"
	"os"

	rundeck "github.com/lusis/go-rundeck/pkg/rundeck.v21"
)

func main() {
	var jobid string
	if len(os.Args) <= 1 {
		fmt.Printf("Usage: rundeck-enable-execution <job uuid>\n")
		os.Exit(1)
	}
	jobid = os.Args[1]
	client, clientErr := rundeck.NewClientFromEnv()
	if clientErr != nil {
		log.Fatal(clientErr.Error())
	}
	err := client.EnableExecution(jobid)
	if err != nil {
		fmt.Printf("%s\n", err)
		os.Exit(1)
	} else {
		fmt.Printf("Execution enabled\n")
		os.Exit(0)
	}
}
