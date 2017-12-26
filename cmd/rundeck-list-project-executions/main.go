package main

import (
	"fmt"
	"log"
	"os"
	"strconv"

	rundeck "github.com/lusis/go-rundeck/pkg/rundeck.v21"
	"github.com/olekukonko/tablewriter"
	kingpin "gopkg.in/alecthomas/kingpin.v2"
)

var (
	projectid = kingpin.Arg("projectid", "").Required().String()
	max       = kingpin.Flag("max", "max number of results to return").Default("200").String()
)

func main() {
	kingpin.Parse()
	client, clientErr := rundeck.NewClientFromEnv()
	if clientErr != nil {
		log.Fatal(clientErr.Error())
	}
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
			if &d.Job != nil {
				name = d.Job.Name
				description = d.Job.Description
			} else {
				name = "<adhoc>"
				description = d.Description
			}
			dateEnded := ""
			if d.DateEnded.Date != nil {
				dateEnded = d.DateEnded.Date.String()
			}
			table.Append([]string{
				strconv.Itoa(d.ID),
				name,
				description,
				d.Status,
				strconv.Itoa(len(d.SuccessfulNodes)) + "/" + strconv.Itoa(len(d.FailedNodes)),
				d.User,
				d.DateStarted.Date.String(),
				dateEnded,
				d.Project,
			})
		}
		table.Render()
	}
}
