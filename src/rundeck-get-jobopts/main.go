package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/olekukonko/tablewriter"
	kingpin "gopkg.in/alecthomas/kingpin.v2"
	rundeck "rundeck.v17"
)

var (
	jobid = kingpin.Arg("jobid", "").Required().String()
)

func main() {
	kingpin.Parse()
	client := rundeck.NewClientFromEnv()
	data, err := client.GetJob(*jobid)
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
