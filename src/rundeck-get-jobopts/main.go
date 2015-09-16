package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/olekukonko/tablewriter"
	rundeck "rundeck.v12"
)

func main() {
	var jobid string
	if len(os.Args) <= 1 {
		fmt.Printf("Usage: rundeck-get-jobopts <job uuid>\n")
		os.Exit(1)
	}
	jobid = os.Args[1]
	client := rundeck.NewClientFromEnv()
	data, err := client.GetJob(jobid)
	if err != nil {
		fmt.Printf("%s\n", err)
	} else {
		scope := data.Job
		table := tablewriter.NewWriter(os.Stdout)
		table.SetHeader([]string{"ID", "Name", "Description", "Options"})
		table.SetAutoWrapText(false)
		var options []string
		for _, d := range *scope.Context.Options {
			var option string
			option = fmt.Sprintf("%s", d.Name)
			if d.Required {
				option = fmt.Sprintf("%s (required)", option)
			}
			options = append(options, option)
		}
		table.Append([]string{scope.ID, scope.Name, scope.Description, strings.Join(options, "\n")})
		table.Render()
	}
}
