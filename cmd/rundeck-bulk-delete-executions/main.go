package main

import (
	"fmt"
	"os"

	rundeck "github.com/lusis/go-rundeck/pkg/rundeck.v19"
	kingpin "gopkg.in/alecthomas/kingpin.v2"
)

var (
	ids = kingpin.Arg("ids", "ids to delete").Required().Strings()
)

func main() {
	kingpin.Parse()
	client := rundeck.NewClientFromEnv()
	res, err := client.DeleteExecutions(*ids)
	if err != nil {
		fmt.Printf("%s\n", err)
		os.Exit(1)
	} else {
		fmt.Printf("Successful: %d\n", res.Successful.Count)
		if res.Failed.Count != 0 {
			fmt.Printf("Failed: %d\n", res.Failed.Count)
			for _, f := range res.Failed.Failures {
				fmt.Printf("%d - %s\n", f.ID, f.Message)
			}
		}
		os.Exit(0)
	}
}
