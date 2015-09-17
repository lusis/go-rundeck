package main

import (
	"fmt"

	kingpin "gopkg.in/alecthomas/kingpin.v2"
	rundeck "rundeck.v12"
)

var (
	jobid  = kingpin.Flag("jobid", "The id of the job to export").Required().String()
	format = kingpin.Flag("format", "Format to export").Default("xml").Enum("xml", "yaml")
)

func main() {
	kingpin.Parse()
	client := rundeck.NewClientFromEnv()
	res := client.ExportJob(*jobid, *format)
	fmt.Printf(res)
}
