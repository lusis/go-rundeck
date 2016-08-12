package main

import (
	"fmt"
	"os"
	"strings"

	kingpin "gopkg.in/alecthomas/kingpin.v2"
	rundeck "rundeck.v17"
)

var (
	runAs      = kingpin.Flag("runas", "user to run as").String()
	nodeFilter = kingpin.Flag("nodeFilter", "node filter to use").String()
	logLevel   = kingpin.Flag("logLevel", "log level to run the job").Enum("DEBUG", "VERBOSE", "INFO", "WARN", "ERROR")
	jobId      = kingpin.Arg("jobId", "Job ID to run").Required().String()
	argString  = kingpin.Arg("argString", "arguments to pass to job").Strings()
)

func main() {
	kingpin.Parse()
	client := rundeck.NewClientFromEnv()
	jobopts := rundeck.RunOptions{
		RunAs:     *runAs,
		LogLevel:  *logLevel,
		Filter:    *nodeFilter,
		Arguments: strings.Join(*argString, " "),
	}
	res, err := client.RunJob(*jobId, jobopts)
	if err != nil {
		fmt.Printf("%s\n", err)
		os.Exit(1)
	} else {
		fmt.Printf("Job %s is %s\n", res.Executions[0].ID, res.Executions[0].Status)
		os.Exit(0)
	}
}
