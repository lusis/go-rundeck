package main

import (
	"fmt"
	"log"
	"os"

	rundeck "github.com/lusis/go-rundeck/pkg/rundeck.v21"
	"github.com/olekukonko/tablewriter"
	kingpin "gopkg.in/alecthomas/kingpin.v2"
)

var (
	formatUsage = fmt.Sprintf("Format to show results [table, csv, list (ids only - useful for piping)]")
	format      = kingpin.Flag("format", formatUsage).Short('F').Default("table").Enum("table", "list", "csv")
	sep         = kingpin.Flag("separator", "separator for csv output").Default(",").String()
	projectid   = kingpin.Arg("projectid", "").Required().String()
)

func main() {
	kingpin.Parse()
	client, clientErr := rundeck.NewClientFromEnv()
	if clientErr != nil {
		log.Fatal(clientErr.Error())
	}
	data, err := client.ListJobs(*projectid)
	if err != nil {
		fmt.Printf("%s\n", err)
	} else {
		if *format == "table" {
			table := tablewriter.NewWriter(os.Stdout)
			table.SetHeader([]string{"ID", "Name", "Description", "Group", "Project"})
			for _, d := range *data {
				table.Append([]string{d.ID, d.Name, d.Description, d.Group, d.Project})
			}
			table.Render()
		} else if *format == "list" {
			for _, d := range *data {
				fmt.Printf("%s\n", d.ID)
			}
		} else if *format == "csv" {
			for _, d := range *data {
				fmt.Printf("%s%s%s%s%s%s%s%s%s\n", d.ID, *sep, d.Name, *sep, d.Description, *sep, d.Group, *sep, d.Project)
			}
		}
		os.Exit(0)
	}
}
