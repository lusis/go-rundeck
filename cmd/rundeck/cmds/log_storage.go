package cmds

import "github.com/spf13/cobra"

func logStorageCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "logstorage",
		Short: "operate on rundeck logstorage",
	}
	cmd.AddCommand(getLogStorageCommand())
	cmd.AddCommand(resumeIncompleteLogStorageCommand())
	return cmd
}
