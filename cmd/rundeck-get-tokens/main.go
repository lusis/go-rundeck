package main

import (
	"fmt"
	"strings"

	cli "github.com/lusis/go-rundeck/pkg/cli"
	"github.com/spf13/cobra"
)

var (
	userid string
)

func runFunc(cmd *cobra.Command, args []string) error {

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
		data, err := cli.Client.GetTokens()
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
		data, err := cli.Client.GetUserTokens(userid)
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
func main() {
	cmd := &cobra.Command{
		Use:   "rundeck-get-tokens [-u user-id]",
		Short: "gets tokens from the rundeck server for the current user or the optionally specified user",
		RunE:  runFunc,
	}
	cmd.Flags().StringVarP(&userid, "user-id", "u", "", "optional user id")
	rootCmd := cli.New(cmd)
	_ = rootCmd.Execute()
}
