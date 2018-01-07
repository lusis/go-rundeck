package cmds

import (
	"github.com/spf13/cobra"
)

const enable = "enable"
const disable = "disable"
const adhoc = "<adhoc>"

// RootCommand is the root of all commands
func RootCommand() {
	cmd := &cobra.Command{
		Use:   "rundeck",
		Short: "Unified rundeck cli binary",
	}
	cmd.AddCommand(projectCommands(),
		listCommands(),
		policiesCommands(),
		jobCommands(),
		jobsCommands(),
		executionCommands(),
		executionsCommands(),
		tokenCommand(),
		tokensCommands(),
		httpCommand(),
		logStorageCommand())
	_ = cmd.Execute()
}
