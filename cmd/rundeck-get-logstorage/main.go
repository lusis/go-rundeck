package main

import (
	"fmt"
	"os"

	rundeck "github.com/lusis/go-rundeck/pkg/rundeck.v19"
	"github.com/olekukonko/tablewriter"
)

func main() {
	client := rundeck.NewClientFromEnv()
	data, err := client.GetLogstorage()
	if err != nil {
		fmt.Printf("%s\n", err)
	} else {
		table := tablewriter.NewWriter(os.Stdout)
		headers := []string{
			"Enabled?",
			"Plugin Name",
			"Succeeded",
			"Failed",
			"Queued",
			"Total",
			"Incomplete",
			"Missing",
		}
		table.SetHeader(headers)
		table.Append([]string{
			fmt.Sprintf("%t", data.Enabled),
			data.PluginName,
			fmt.Sprintf("%d", data.SucceededCount),
			fmt.Sprintf("%d", data.FailedCount),
			fmt.Sprintf("%d", data.QueuedCount),
			fmt.Sprintf("%d", data.TotalCount),
			fmt.Sprintf("%d", data.IncompleteCount),
			fmt.Sprintf("%d", data.MissingCount),
		})
		table.Render()
	}
}
