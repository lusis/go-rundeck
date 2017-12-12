package main

import (
	"fmt"
	"os"
	"strings"

	rundeck "github.com/lusis/go-rundeck/pkg/rundeck.v19"
	"github.com/olekukonko/tablewriter"
	kingpin "gopkg.in/alecthomas/kingpin.v2"
)

var (
	projectid = kingpin.Arg("projectname", "").Required().String()
	jobid     = kingpin.Arg("jobname", "").Required().Strings()
)

func main() {
	kingpin.Parse()
	client := rundeck.NewClientFromEnv()
	data, err := client.FindJobByName(strings.Join(*jobid, " "), *projectid)
	if err != nil {
		fmt.Printf("%s\n", err)
	} else {
		scope := *data
		table := tablewriter.NewWriter(os.Stdout)
		table.SetHeader([]string{"ID", "Name", "Description", "Group", "Steps", "Node Filters"})
		var steps []string
		var nodefilters []string
		for _, d := range scope.Sequence.Steps {
			var stepDescription string
			if d.Description == "" {
				if d.JobRef != nil {
					stepDescription = d.JobRef.Name
				} else if d.Exec != nil {
					stepDescription = *d.Exec
				}
			} else {
				stepDescription = d.Description
			}
			steps = append(steps, stepDescription)
		}
		nodefilters = append(nodefilters, scope.NodeFilters.Filter...)
		table.Append([]string{scope.ID, scope.Name, scope.Description, scope.Group, strings.Join(steps, "\n"), strings.Join(nodefilters, "\n")})
		table.Render()
	}
}
