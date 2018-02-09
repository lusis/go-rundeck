package cmds

import (
	"bufio"
	"os"
	"strconv"
	"time"

	"github.com/lusis/go-rundeck/pkg/cli"
	"github.com/spf13/cobra"
)

var tailPollInterval int

func getOutput(eID int, offset string, w *bufio.Writer) error {
	wait := time.Duration(tailPollInterval)
	time.Sleep(wait * time.Second)
	iOffset, iOffsetErr := strconv.Atoi(offset)
	if iOffsetErr != nil {
		return iOffsetErr
	}
	data, err := cli.Client.GetExecutionOutputWithOffset(eID, iOffset)
	if err != nil {
		return err
	}
	for _, entry := range data.Entries {
		_, err := w.WriteString(entry.Log + "\n")
		if err != nil {
			return err
		}

		flushErr := w.Flush()
		if flushErr != nil {
			return flushErr
		}
	}

	if data.ExecCompleted {
		return nil
	}

	return getOutput(eID, data.Offset, w)
}

func tailExecutionOutputFunc(cmd *cobra.Command, args []string) error {
	id := args[0]
	eID, eIDerr := strconv.Atoi(id)
	if eIDerr != nil {
		return eIDerr
	}
	w := bufio.NewWriter(os.Stdout)
	return getOutput(eID, "0", w)
}

func tailExecutionOutputCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "tail execution-id [-i poll-interval]",
		Short: "tails an execution's output from the rundeck server",
		Args:  cobra.MinimumNArgs(1),
		RunE:  tailExecutionOutputFunc,
	}
	rootCmd := cli.New(cmd)
	rootCmd.ResetFlags()
	rootCmd.Flags().IntVarP(&tailPollInterval, "interval", "i", 2, "interval to poll for more log data in seconds")
	return rootCmd
}
