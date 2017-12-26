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
	jobid = kingpin.Arg("jobid", "").Required().String()
)

func main() {
	kingpin.Parse()
	client, clientErr := rundeck.NewClientFromEnv()
	if clientErr != nil {
		log.Fatal(clientErr.Error())
	}
	data, err := client.GetJobOpts(*jobid)
	if err != nil {
		fmt.Printf("%s\n", err)
		os.Exit(1)
	} else {
		table := tablewriter.NewWriter(os.Stdout)
		table.SetHeader([]string{
			"Name",
			"Description",
			"Value",
			"Required?",
			"Regex",
		})
		table.SetAutoWrapText(false)
		for _, d := range data {
			table.Append([]string{
				d.Name,
				d.Description,
				d.Value,
				fmt.Sprintf("%t", d.Required),
				d.Regex})
		}
		table.Render()
	}
}
