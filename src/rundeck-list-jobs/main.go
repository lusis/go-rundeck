package main

import (
	"fmt"
	"os"

	"github.com/olekukonko/tablewriter"
	kingpin "gopkg.in/alecthomas/kingpin.v2"
	rundeck "rundeck.v13"
)

var (
	projectid = kingpin.Arg("projectid", "").Required().String()
)

func main() {
	kingpin.Parse()
	client := rundeck.NewClientFromEnv()
	data, err := client.ListJobs(*projectid)
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
