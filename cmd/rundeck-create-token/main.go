package main

import (
	"fmt"
	"strings"

	cli "github.com/lusis/go-rundeck/pkg/cli"
	"github.com/spf13/cobra"
)

func runFunc(cmd *cobra.Command, args []string) error {
	user := args[0]
	token, err := cli.Client.CreateToken(user)
	if err != nil {
		return err
	}
	cli.OutputFormatter.SetHeaders([]string{
		"ID",
		"User",
		"Token",
		"Creator",
		"Expiration",
		"Roles",
		"Expired",
		"Duration",
	})
	rowErr := cli.OutputFormatter.AddRow([]string{
		token.ID,
		token.User,
		token.Creator,
		token.Expiration.String(),
		strings.Join(token.Roles, ", "),
		fmt.Sprintf("%t", token.Expired),
		token.Duration,
	})
	if rowErr != nil {
		return rowErr
	}
	cli.OutputFormatter.Draw()
	return nil
}

func main() {
	cmd := &cobra.Command{
		Use:   "rundeck-create-token username",
		Short: "creates an api token in rundeck for the named user",
		Args:  cobra.MinimumNArgs(1),
		RunE:  runFunc,
	}
	rootCmd := cli.New(cmd)
	_ = rootCmd.Execute()
}
