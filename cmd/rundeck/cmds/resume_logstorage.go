package cmds

import (
	"fmt"

	cli "github.com/lusis/go-rundeck/pkg/cli"
	"github.com/spf13/cobra"
)

func resumeIncompleteLogStorageFunc(cmd *cobra.Command, args []string) error {
	data, err := cli.Client.ResumeIncompleteLogStorage()
	if err != nil {
		return err
	}
	fmt.Printf("%t\n", data)
	return nil
}

func resumeIncompleteLogStorageCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "resume",
		Short: "resumes incomplete log storage processing on the rundeck server",
		RunE:  resumeIncompleteLogStorageFunc,
	}
	rootCmd := cli.New(cmd)
	return rootCmd
}
