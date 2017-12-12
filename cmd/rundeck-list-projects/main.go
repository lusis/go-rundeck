package main

import (
	"fmt"
	"os"

	"github.com/olekukonko/tablewriter"
	kingpin "gopkg.in/alecthomas/kingpin.v2"

	rundeck "github.com/lusis/go-rundeck/pkg/rundeck.v19"
)

var (
	formatUsage = fmt.Sprintf("Format to show results [table, csv, list (ids only - useful for piping)]")
	format      = kingpin.Flag("format", formatUsage).Short('F').Default("table").Enum("table", "list", "csv")
	sep         = kingpin.Flag("separator", "separator for csv output").Default(",").String()
	header      = kingpin.Flag("headers", "add headers for csv output").Default("false").Bool()
)

func main() {
	kingpin.Parse()
	client := rundeck.NewClientFromEnv()
	data, err := client.ListProjects()
	if err != nil {
		fmt.Printf("%s\n", err)
		os.Exit(1)
	} else {
		if *format == "table" {
			table := tablewriter.NewWriter(os.Stdout)
			table.SetHeader([]string{
				"Name",
				"Description",
				"URL",
			})
			for _, d := range data.Projects {
				table.Append([]string{
					d.Name,
					d.Description,
					d.Url,
				})
			}
			table.Render()
		} else if *format == "list" {
			for _, d := range data.Projects {
				fmt.Printf("%s\n", d.Name)
			}
		} else if *format == "csv" {
			if *header == true {
				fmt.Printf("%s%s%s%s%s\n", "NAME", *sep, "DESCRIPTION", *sep, "URL")
			}
			for _, d := range data.Projects {
				fmt.Printf("%s%s%s%s%s\n", d.Name, *sep, d.Description, *sep, d.Url)
			}
		} else {
			fmt.Printf("Unknown output format: %s\n", *format)
			os.Exit(1)
		}
		os.Exit(0)
	}
}
