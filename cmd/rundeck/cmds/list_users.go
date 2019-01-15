package cmds

import (
	"fmt"

	"github.com/lusis/go-rundeck/pkg/cli"
	"github.com/spf13/cobra"
)

func listUsersFunc(cmd *cobra.Command, args []string) error {
	data, err := cli.Client.ListUsers()
	if err != nil {
		return err
	}
	cli.OutputFormatter.SetHeaders([]string{"Login", "First Name", "Last Name", "Email", "Created", "Updated", "Last Job", "Tokens"})
	for _, d := range data {
		created, updated, lastjob := d.Created.Format(cli.TimeFormat), d.Updated.Format(cli.TimeFormat), d.LastJob.Format(cli.TimeFormat)
		if err := cli.OutputFormatter.AddRow([]string{d.Login, d.FirstName, d.LastName, d.Email, created, updated, lastjob, fmt.Sprintf("%d", d.Tokens)}); err != nil {
			return err
		}
	}
	cli.OutputFormatter.Draw()
	return nil
}

func listUsersCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "users",
		Short: "gets a list of users from the rundeck server",
		RunE:  listUsersFunc,
	}
	rootCmd := cli.New(cmd)
	return rootCmd
}
