package main

import (
	"fmt"
	"os"

	kingpin "gopkg.in/alecthomas/kingpin.v2"
	rundeck "rundeck.v17"
)

var (
	filename = kingpin.Flag("file", "Full /path/to/file to import").Required().ExistingFile()
	format   = kingpin.Flag("format", "Format to import").Default("xml").Enum("xml", "yaml")
	dupe     = kingpin.Flag("dupe", "How to handle existing jobs with same name").Default("create").Enum("create", "update", "skip")
	uuid     = kingpin.Flag("uuid", "Preserve or strip uuids").Default("preserve").Enum("preserve", "remove")
	project  = kingpin.Flag("project", "Project name for imported job").Required().String()
)

func main() {
	kingpin.Parse()
	client := rundeck.NewClientFromEnv()
	importParams := rundeck.ImportParams{
		Filename: *filename,
		Format:   *format,
		Dupe:     *dupe,
		Uuid:     *uuid,
		Project:  *project,
	}
	jobid, err := client.ImportJob(importParams)
	if err != nil {
		fmt.Printf("%s\n", err)
		os.Exit(1)
	} else {
		fmt.Printf("Job %s imported\n", jobid)
		os.Exit(0)
	}
}
