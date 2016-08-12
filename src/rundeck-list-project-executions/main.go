package main

import (
	"fmt"
	"os"
	"strconv"

	"github.com/olekukonko/tablewriter"
	kingpin "gopkg.in/alecthomas/kingpin.v2"
	rundeck "rundeck.v17"
)

var (
	projectid = kingpin.Arg("projectid", "").Required().String()
	max       = kingpin.Flag("max", "max number of results to return").Default("200").String()
)

func main() {
	kingpin.Parse()
	client := rundeck.NewClientFromEnv()
	options := make(map[string]string)
	options["max"] = *max
	data, err := client.ListProjectExecutions(*projectid, options)
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
				d.DateStarted,
				d.DateEnded,
				d.Project,
			})
		}
		table.Render()
	}
}
