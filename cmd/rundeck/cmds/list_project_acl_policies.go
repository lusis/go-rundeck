package cmds

import (
	"github.com/lusis/go-rundeck/pkg/cli"
	"github.com/spf13/cobra"
)

func listProjectACLPoliciesFunc(cmd *cobra.Command, args []string) error {
	projectName := args[0]
	policies, err := cli.Client.ListProjectACLPolicies(projectName)
	if err != nil {
		return err
	}
	cli.OutputFormatter.SetHeaders([]string{
		"Name",
		"Path",
		"Type",
		"HRef",
		"Parent",
		"Parent Type",
	})
	parent := "/"
	if policies.Path != "" {
		parent = policies.Path
	}
	for _, p := range policies.Resources {
		if err := cli.OutputFormatter.AddRow([]string{
			p.Name,
			p.Path,
			p.Type,
			p.Href,
			parent,
			policies.Type,
		}); err != nil {
			return err
		}
	}
	cli.OutputFormatter.Draw()
	return nil
}

func listProjectACLPoliciesCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "list",
		Short: "lists system acl policies",
		Args:  cobra.MinimumNArgs(1),
		RunE:  listProjectACLPoliciesFunc,
	}
	rootCmd := cli.New(cmd)
	return rootCmd
}
