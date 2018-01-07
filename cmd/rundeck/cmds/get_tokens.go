package cmds

import (
	"fmt"
	"strings"

	"github.com/lusis/go-rundeck/pkg/cli"
	"github.com/spf13/cobra"
)

var (
	getTokensCmd = &cobra.Command{
		Use:   "get username",
		Short: "gets all tokens for the specified username",
		RunE:  getTokensFunc,
		Args:  cobra.MinimumNArgs(1),
	}
)

func getTokensFunc(cmd *cobra.Command, args []string) error {
	userid := args[0]
	cli.OutputFormatter.SetHeaders([]string{
		"ID",
		"User",
		"Creator",
		"Duration",
		"Expiration",
		"Expired?",
		"Roles",
	})
	if userid == "" {
		data, err := cli.Client.ListTokens()
		if err != nil {
			return err
		}
		for _, d := range data {
			if rowErr := cli.OutputFormatter.AddRow([]string{
				d.ID,
				d.User,
				d.Creator,
				d.Duration,
				d.Expiration.String(),
				fmt.Sprintf("%t", d.Expired),
				strings.Join(d.Roles, ","),
			}); rowErr != nil {
				return rowErr
			}
		}
	} else {
		data, err := cli.Client.ListTokensForUser(userid)
		if err != nil {
			return err
		}
		for _, d := range data {
			if rowErr := cli.OutputFormatter.AddRow([]string{
				d.ID,
				d.User,
				d.Creator,
				d.Duration,
				d.Expiration.String(),
				fmt.Sprintf("%t", d.Expired),
				strings.Join(d.Roles, ","),
			}); rowErr != nil {
				return rowErr
			}
		}
	}
	cli.OutputFormatter.Draw()
	return nil
}
func getTokensCommand() *cobra.Command {
	cmd := cli.New(getTokensCmd)
	return cmd
}
