package main

import (
	"fmt"
	"os"

	rundeck "github.com/lusis/go-rundeck/pkg/rundeck.v19"
	"github.com/olekukonko/tablewriter"
)

func main() {
	client := rundeck.NewClientFromEnv()
	data, err := client.GetTokens()
	if err != nil {
		fmt.Printf("%s\n", err)
	} else {
		table := tablewriter.NewWriter(os.Stdout)
		table.SetHeader([]string{
			"ID",
			"User",
		})
		for _, d := range data.Tokens {
			table.Append([]string{
				d.ID,
				d.User,
			})
		}
		table.Render()
	}
}
