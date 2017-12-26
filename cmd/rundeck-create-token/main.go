package main

import (
	"errors"
	"fmt"
	"strings"

	cli "github.com/lusis/go-rundeck/pkg/cli"
	"github.com/spf13/cobra"
)

var (
	user string
)

func runFunc(cmd *cobra.Command, args []string) error {
	if user == "" {
		return errors.New("username is required")
	}
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
		Use:   "rundeck-create-token -u username",
		Short: "creates an api token in rundeck for the named user",
		RunE:  runFunc,
	}
	cmd.Flags().StringVarP(&user, "user", "u", "", "username")
	rootCmd := cli.New(cmd)
	_ = rootCmd.Execute()
}
