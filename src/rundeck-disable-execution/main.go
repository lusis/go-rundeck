package main

import (
	"fmt"
	"os"

	rundeck "rundeck.v17"
)

func main() {
	var jobid string
	if len(os.Args) <= 1 {
		fmt.Printf("Usage: rundeck-disable-execution <job uuid>\n")
		os.Exit(1)
	}
	jobid = os.Args[1]
	client := rundeck.NewClientFromEnv()
	err := client.DisableExecution(jobid)
	if err != nil {
		fmt.Printf("%s\n", err)
		os.Exit(1)
	} else {
		fmt.Printf("Execution disabled\n")
		os.Exit(0)
	}
}
