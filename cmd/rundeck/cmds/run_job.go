package cmds

import (
	"strconv"
	"time"

	"github.com/lusis/go-rundeck/pkg/cli"
	rundeck "github.com/lusis/go-rundeck/pkg/rundeck.v21"
	"github.com/spf13/cobra"
)

var (
	runJobRunAtTime  string
	runJobRunAs      string
	runJobNodeFilter string
	runJobLogLevel   string
	runJobArgString  string
	runJobTimeFormat string
	runJobOptions    []string
)

const runJobDefaultTimeFormat = "2006-01-02T15:04:05-0700"

func runJobFunc(cmd *cobra.Command, args []string) error {
	jobid := args[0]
	var runOpts []rundeck.RunJobOption
	params, paramErr := cli.BuildParams(runJobOptions)
	if paramErr != nil {
		return paramErr
	}
	if len(params) > 0 {
		runOpts = append(runOpts, rundeck.RunJobOpts(params))
	}

	if runJobRunAs != "" {
		runOpts = append(runOpts, rundeck.RunJobAs(runJobRunAs))
	}
	if runJobNodeFilter != "" {
		runOpts = append(runOpts, rundeck.RunJobFilter(runJobNodeFilter))
	}
	if runJobArgString != "" {
		runOpts = append(runOpts, rundeck.RunJobArgs(runJobArgString))
	}
	if runJobLogLevel != "" {
		runOpts = append(runOpts, rundeck.RunJobLogLevel(runJobLogLevel)) // nolint: ineffassign
	}
	if runJobRunAtTime != "" {
		format := runJobDefaultTimeFormat
		if runJobTimeFormat != "" {
			format = runJobTimeFormat
		}
		rt, rtErr := time.Parse(format, runJobRunAtTime)
		if rtErr != nil {
			return rtErr
		}
		runOpts = append(runOpts, rundeck.RunJobRunAt(rt))
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
		name = adhoc
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
func runJobCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "run job-id [-q foo=bar] [-c application/json]",
		Short: "runs a rundeck job",
		Long:  runJobLongHelp,
		RunE:  runJobFunc,
		Args:  cobra.MinimumNArgs(1),
	}
	rootCmd := cli.New(cmd)
	rootCmd.Flags().StringSliceVarP(&runJobOptions, "options", "o", []string{}, "custom options to pass in format of name=value. Can specify multiple times")
	rootCmd.Flags().StringVarP(&runJobRunAs, "user", "u", "", "user to run as")
	rootCmd.Flags().StringVarP(&runJobNodeFilter, "filter", "f", "", "node filter to use")
	rootCmd.Flags().StringVarP(&runJobArgString, "argstring", "a", "", "args string to use")
	rootCmd.Flags().StringVarP(&runJobLogLevel, "loglevel", "l", "", "log level to use")
	rootCmd.Flags().StringVarP(&runJobRunAtTime, "time", "t", "", "when to run the job. If no format is specified "+runJobDefaultTimeFormat+" is used")
	rootCmd.Flags().StringVar(&runJobTimeFormat, "time-format", runJobDefaultTimeFormat, "golang time format string")

	return rootCmd
}

const runJobLongHelp = `
# Simple instant execution
rundeck job run <job-id>

# Run at a specified time using the default time format
rundeck job run <job-id> -t 2018-01-03T11:30:00-0500

# Run on the same job above but using a custom time format
rundeck job run <job-id> -t 2018-01-03 --time-format=2006-01-02

# Run a job with a custom arg string (you need to quote the full argstring)
rundeck job run <job-id> -a "-sleeptime 30"

# Run the job above but using options instead of argstring
rundeck job run <job-id> -o sleeptime=30 -o someparam=anotherval

# Run as another user
rundeck job run <job-id> -u another-user
`
