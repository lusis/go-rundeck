package main

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
	"time"

	cli "github.com/lusis/go-rundeck/pkg/cli"
	rundeck "github.com/lusis/go-rundeck/pkg/rundeck.v21"
	"github.com/spf13/cobra"
)

var (
	runAtTime  string
	runAs      string
	nodeFilter string
	logLevel   string
	argString  string
	timeFormat string
	options    []string
)

const defaultTimeFormat = "2006-01-02T15:04:05-0700"

func buildParams(values []string) (map[string]string, error) {
	p := map[string]string{}
	for _, value := range values {
		parts := strings.SplitN(value, "=", 2)
		if len(parts) != 2 {
			return nil, fmt.Errorf("expected key=value got '%s'", value)
		}
		if parts[1] == "" {
			msg := fmt.Sprintf("missing value for key '%s'", parts[0])
			return nil, errors.New(msg)
		}
		if parts[0] == "" {
			msg := fmt.Sprintf("missing key for value '%s'", parts[1])
			return nil, errors.New(msg)
		}
		p[parts[0]] = parts[1]
	}
	return p, nil
}

func runFunc(cmd *cobra.Command, args []string) error {
	jobid := args[0]
	var runOpts []rundeck.RunJobOption
	params, paramErr := buildParams(options)
	if paramErr != nil {
		return paramErr
	}
	if len(params) > 0 {
		runOpts = append(runOpts, rundeck.RunJobOpts(params))
	}

	if runAs != "" {
		runOpts = append(runOpts, rundeck.RunJobAs(runAs))
	}
	if nodeFilter != "" {
		runOpts = append(runOpts, rundeck.RunJobFilter(nodeFilter))
	}
	if argString != "" {
		runOpts = append(runOpts, rundeck.RunJobArgs(argString))
	}
	if logLevel != "" {
		runOpts = append(runOpts, rundeck.RunJobLogLevel(logLevel)) // nolint: ineffassign
	}
	if runAtTime != "" {
		format := defaultTimeFormat
		if timeFormat != "" {
			format = timeFormat
		}
		rt, rtErr := time.Parse(format, runAtTime)
		if rtErr != nil {
			return rtErr
		}
		runOpts = append(runOpts, rundeck.RunJobAt(rt))
	}
	data, err := cli.Client.RunJob(jobid, runOpts...)
	if err != nil {
		return err
	}
	cli.OutputFormatter.SetHeaders([]string{
		"ID",
		"Job Name",
		"Job Description",
		"Arguments",
		"Node Success/Failure Count",
		"User",
		"Project",
	})

	var description string
	var name string
	if &data.Job != nil {
		name = data.Job.Name
		description = data.Job.Description
	} else {
		name = "<adhoc>"
		description = data.Description
	}
	if rowErr := cli.OutputFormatter.AddRow([]string{
		strconv.Itoa(data.ID),
		name,
		description,
		data.ArgString,
		strconv.Itoa(len(data.SuccessfulNodes)) + "/" + strconv.Itoa(len(data.FailedNodes)),
		data.User,
		data.Project,
	}); rowErr != nil {
		return rowErr
	}
	cli.OutputFormatter.Draw()
	return nil
}

func main() {
	cmd := &cobra.Command{
		Use:   "rundeck-runjob job-id [-q foo=bar] [-c application/json]",
		Short: "runs a rundeck job",
		Long:  longHelp,
		RunE:  runFunc,
		Args:  cobra.MinimumNArgs(1),
	}
	cmd.Flags().StringSliceVarP(&options, "options", "o", []string{}, "custom options to pass in format of name=value. Can specify multiple times")
	cmd.Flags().StringVarP(&runAs, "user", "u", "", "user to run as")
	cmd.Flags().StringVarP(&nodeFilter, "filter", "f", "", "node filter to use")
	cmd.Flags().StringVarP(&argString, "argstring", "a", "", "args string to use")
	cmd.Flags().StringVarP(&logLevel, "loglevel", "l", "", "log level to use")
	cmd.Flags().StringVarP(&runAtTime, "time", "t", "", "when to run the job. If no format is specified "+defaultTimeFormat+" is used")
	cmd.Flags().StringVar(&timeFormat, "time-format", defaultTimeFormat, "golang time format string")
	rootCmd := cli.New(cmd)
	_ = rootCmd.Execute()
}

const longHelp = `
# Simple instant execution
rundeck-run-job <job-id>

# Run at a specified time using the default time format
rundeck-run-job <job-id> -t 2018-01-03T11:30:00-0500

# Run on the same job above but using a custom time format
rundeck-run-job <job-id> -t 2018-01-03 --time-format=2006-01-02

# Run a job with a custom arg string (you need to quote the full argstring)
rundeck-run-job <job-id> -a "-sleeptime 30"

# Run the job above but using options instead of argstring
rundeck-run-job <job-id> -o sleeptime=30 -o someparam=anotherval

# Run as another user
rundeck-run-job <job-id> -u another-user
`
