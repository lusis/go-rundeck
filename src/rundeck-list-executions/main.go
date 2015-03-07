package main

import (
	"fmt"
	"os"
	"strconv"

	"github.com/olekukonko/tablewriter"
	rundeck "rundeck.v12"
)

func main() {
	var projectid string
	if len(os.Args) <= 1 {
		fmt.Printf("Usage: rundeck-list-executions <project id>\n")
		os.Exit(1)
	}
	projectid = os.Args[1]
	client := rundeck.NewClientFromEnv()
	options := make(map[string]string)
	options["max"] = "200"
	data, err := client.ListExecutions(projectid, options)
	if err != nil {
		fmt.Printf("%s\n", err)
	} else {
		table := tablewriter.NewWriter(os.Stdout)
		table.SetHeader([]string{
			"ID",
			"Job Name",
			"Job Description",
			"Status",
			"Node Success/Failure Count",
			"User",
			"Start",
			"End",
			"Project",
		})
		for _, d := range data.Executions {
			var description string
			var name string
			if d.Job != nil {
				name = d.Job.Name
				description = d.Job.Description
			} else {
				name = "<adhoc>"
				description = d.Description
			}
			table.Append([]string{
				d.ID,
				name,
				description,
				d.Status,
				strconv.Itoa(len(d.SuccessfulNodes.Nodes)) + "/" + strconv.Itoa(len(d.FailedNodes.Nodes)),
				d.User,
				strconv.FormatInt(d.DateStarted.UnixTime, 10),
				strconv.FormatInt(d.DateEnded.UnixTime, 10),
				d.Project,
			})
		}
		table.Render()
	}
}
