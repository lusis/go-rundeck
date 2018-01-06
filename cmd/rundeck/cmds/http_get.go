package cmds

import (
	"fmt"

	cli "github.com/lusis/go-rundeck/pkg/cli"
	httpclient "github.com/lusis/go-rundeck/pkg/httpclient"
	"github.com/spf13/cobra"
)

var (
	httpGetQueryParameters []string
	httpGetContentType     string
)

var helpLong = `
"rundeck http get" allows you to perform authenticated http get requests against the rundeck api.
Examples:

# rundeck http get system/acl/
# rundeck http get job/XXXXXXX/executions -q max=1
# rundeck http get job/XXXXXXX/executions -q max=1 -q status=failed
# rundeck http get execution/29 -c application/xml

This tool is used to generate test response data for this library itself.
`

func httpGetFunc(cmd *cobra.Command, args []string) error {
	path := args[0]
	params, paramErr := cli.BuildParams(httpGetQueryParameters)
	if paramErr != nil {
		return paramErr
	}
	res, err := cli.Client.Get(path, httpclient.QueryParams(params), httpclient.Accept(httpGetContentType))
	if err != nil {
		return err
	}
	fmt.Println(string(res))
	return nil
}

func httpGetCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "get path [-q foo=bar] [-c application/json]",
		Short: "performs an authenticated http get against the rundeck api",
		Long:  helpLong,
		RunE:  httpGetFunc,
		Args:  cobra.MinimumNArgs(1),
	}
	rootCmd := cli.New(cmd)
	rootCmd.ResetFlags() // remove the global --output-format flag inherited from cli.New()
	rootCmd.Flags().StringSliceVarP(&httpGetQueryParameters, "query-param", "q", []string{}, "custom query params to pass in format of name=value. Can specify multiple times")
	rootCmd.Flags().StringVarP(&httpGetContentType, "content-type", "c", "application/json", "content-type to return")

	return rootCmd
}
