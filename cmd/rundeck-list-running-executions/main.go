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
	project     = kingpin.Arg("project", "").Required().String()
	formatUsage = fmt.Sprintf("Format to show results [table, csv, list (ids only - useful for piping)]")
	format      = kingpin.Flag("format", formatUsage).Short('F').Default("table").Enum("table", "list", "csv")
	sep         = kingpin.Flag("separator", "separator for csv output").Default(",").String()
	header      = kingpin.Flag("headers", "add headers for csv output").Default("false").Bool()
)

func main() {
	kingpin.Parse()
	client, clientErr := rundeck.NewClientFromEnv()
	if clientErr != nil {
		log.Fatal(clientErr.Error())
	}
	res, err := client.ListRunningExecutions(*project)
	if err != nil {
		fmt.Printf("%s\n", err)
		os.Exit(1)
	} else {
		if *format == "table" {
			table := tablewriter.NewWriter(os.Stdout)
			table.SetHeader([]string{
				"ID",
				"Status",
				"User",
			})
			for _, d := range res.Executions {
				table.Append([]string{
					strconv.Itoa(d.ID),
					d.Status,
					d.User,
				})
			}
			table.Render()
		} else if *format == "csv" {
			if *header {
				fmt.Printf("ID%sStatus%sUser", *sep, *sep)
			}
			for _, d := range res.Executions {
				fmt.Printf("%d%s%s%s%s\n", d.ID, *sep, d.Status, *sep, d.User)
			}
		} else if *format == "list" {
			for _, d := range res.Executions {
				fmt.Printf("%d\n", d.ID)
			}
		} else {
			fmt.Printf("Unknown output format: %s\n", *format)
			os.Exit(1)
		}
		os.Exit(0)
	}
}
