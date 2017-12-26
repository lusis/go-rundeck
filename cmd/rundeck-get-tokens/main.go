package main

import (
	"fmt"
	"log"
	"os"
	"strings"

	rundeck "github.com/lusis/go-rundeck/pkg/rundeck.v21"
	"github.com/olekukonko/tablewriter"
)

func main() {
	client, clientErr := rundeck.NewClientFromEnv()
	if clientErr != nil {
		log.Fatal(clientErr.Error())
	}
	data, err := client.GetTokens()
	if err != nil {
		fmt.Printf("%s\n", err)
	} else {
		table := tablewriter.NewWriter(os.Stdout)
		table.SetHeader([]string{
			"ID",
			"User",
			"Creator",
			"Duration",
			"Expiration",
			"Expired?",
			"Roles",
		})
		for _, d := range data {
			table.Append([]string{
				d.ID,
				d.User,
				d.Creator,
				d.Duration,
				d.Expiration.String(),
				fmt.Sprintf("%t", d.Expired),
				strings.Join(d.Roles, ","),
			})
		}
		table.Render()
	}
}
