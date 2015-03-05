package main

import (
	"fmt"
	"os"

	"github.com/olekukonko/tablewriter"
	rundeck "rundeck.v12"
)

func main() {
	var projectid string
	if len(os.Args) <= 1 {
		fmt.Printf("Usage: rundeck-get-job <project id>\n")
		os.Exit(1)
	}
	projectid = os.Args[1]
	client := rundeck.NewClientFromEnv()
	data, err := client.ListJobs(projectid)
	if err != nil {
		fmt.Printf("%s\n", err)
	} else {
		table := tablewriter.NewWriter(os.Stdout)
		table.SetHeader([]string{"ID", "Name", "Description", "Group", "Project"})
		for _, d := range data.Jobs {
			table.Append([]string{d.ID, d.Name, d.Description, d.Group, d.Project})
		}
		table.Render()
	}
}
