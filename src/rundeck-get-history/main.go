package main

import (
	"fmt"
	"os"

	"github.com/olekukonko/tablewriter"
	kingpin "gopkg.in/alecthomas/kingpin.v2"

	rundeck "rundeck.v17"
)

var (
	projectid = kingpin.Arg("projectid", "").Required().String()
)

func main() {
	kingpin.Parse()
	client := rundeck.NewClientFromEnv()
	top, err := client.GetHistory(*projectid)
	if err != nil {
		fmt.Printf("%s\n", err)
	} else {
		table := tablewriter.NewWriter(os.Stdout)
		for _, data := range top.Events {
			var job string
			if data.Job != nil {
				job = data.Job.ID
			} else {
				job = "<adhoc>"
			}
			table.SetHeader([]string{"Status", "Summary", "Start Time", "End Time", "S/F/T", "Job", "Execution", "User", "Project"})
			table.Append([]string{
				data.Status,
				data.Summary,
				data.StartTime,
				data.EndTime,
				fmt.Sprintf("%d/%d/%d", data.NodeSummary.Succeeded, data.NodeSummary.Failed, data.NodeSummary.Total),
				job,
				fmt.Sprintf("%d", data.Execution.ID),
				data.User,
				data.Project,
			})
		}
		table.Render()
	}
}
