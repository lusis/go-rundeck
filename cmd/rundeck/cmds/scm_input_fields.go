package cmds

import (
	"fmt"
	"strings"

	"github.com/lusis/go-rundeck/pkg/cli"
	"github.com/spf13/cobra"
)

func getSCMPluginInputFieldsFunc(cmd *cobra.Command, args []string) error {
	projectName := args[0]
	integration := args[1]
	pluginType := args[2]
	data, err := cli.Client.GetSCMPluginInputFields(projectName, integration, pluginType)
	if err != nil {
		return err
	}
	cli.OutputFormatter.SetHeaders([]string{
		"Name",
		"Title",
		"Type",
		"Default Value",
		"Description",
		"Required?",
		"Values",
	})
	for _, p := range data.Fields {
		if scmInputFieldsRequiredOnly {
			if !p.Required {
				continue
			}
		}
		// by default we only want the first line of the description
		desc := strings.Split(p.Description, "\n")[0]

		if scmInputFieldsFullDescription {
			desc = p.Description
		}
		if rowErr := cli.OutputFormatter.AddRow([]string{
			p.Name,
			p.Title,
			p.Type,
			p.DefaultValue,
			desc,
			fmt.Sprintf("%t", p.Required),
			strings.Join(p.Values, ","),
		}); rowErr != nil {
			return err
		}
	}
	cli.OutputFormatter.Draw()
	return nil
}

func getSCMPluginInputFieldsCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "input-fields project-name integration type [--full-description] [--required-only]",
		Short: "lists the input fields for a given SCM plugin",
		Args:  cobra.MinimumNArgs(3),
		RunE:  getSCMPluginInputFieldsFunc,
	}
	rootCmd := cli.New(cmd)
	rootCmd.Flags().BoolVar(&scmInputFieldsFullDescription, "full-description", false, "Show full field description")
	rootCmd.Flags().BoolVarP(&scmInputFieldsRequiredOnly, "required-only", "r", false, "Show only required fields")
	return rootCmd
}
