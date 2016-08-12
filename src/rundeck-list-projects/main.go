package main

import (
	"fmt"
	"os"

	"github.com/olekukonko/tablewriter"

	rundeck "rundeck.v17"
)

func main() {
	client := rundeck.NewClientFromEnv()
	data, err := client.ListProjects()
	if err != nil {
		fmt.Printf("%s\n", err)
	} else {
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
	}
}
