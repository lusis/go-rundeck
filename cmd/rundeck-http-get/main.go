package main

import (
	"errors"
	"fmt"
	"strings"

	cli "github.com/lusis/go-rundeck/pkg/cli"
	httpclient "github.com/lusis/go-rundeck/pkg/httpclient"
	"github.com/spf13/cobra"
)

var (
	queryParameters []string
	contentType     string
)

var helpLong = `
rundeck-http-get allows you to perform authenticated http get requests against the rundeck api.
Examples:

# rundeck-http-get system/acl/
# rundeck-http-get job/XXXXXXX/executions -q max=1
# rundeck-http-get job/XXXXXXX/executions -q max=1 -q status=failed
# rundeck-http-get execution/29 -c application/xml

This tool is used to generate test response data for this library itself.
`

func buildParams(values []string) (map[string]string, error) {
	p := map[string]string{}
	for _, value := range values {
		parts := strings.SplitN(value, "=", 2)
		if len(parts) != 2 {
			return nil, fmt.Errorf("expected key=value got '%s'", value)
		}
		if parts[1] == "" {
			msg := fmt.Sprintf("missing value for key '%s'", parts[0])
			return nil, errors.New(msg)
		}
		if parts[0] == "" {
			msg := fmt.Sprintf("missing key for value '%s'", parts[1])
			return nil, errors.New(msg)
		}
		p[parts[0]] = parts[1]
	}
	return p, nil
}

func runFunc(cmd *cobra.Command, args []string) error {
	path := args[0]
	params, paramErr := buildParams(queryParameters)
	if paramErr != nil {
		return paramErr
	}
	res, err := cli.Client.Get(path, httpclient.QueryParams(params), httpclient.Accept(contentType))
	if err != nil {
		return err
	}
	fmt.Println(string(res))
	return nil
}

func main() {
	cmd := &cobra.Command{
		Use:   "rundeck-http-get path [-q foo=bar] [-c application/json]",
		Short: "performs an authenticated http get against the rundeck api",
		Long:  helpLong,
		RunE:  runFunc,
		Args:  cobra.MinimumNArgs(1),
	}
	cli.UseFormatter = false
	cmd.Flags().StringSliceVarP(&queryParameters, "query-param", "q", []string{}, "custom query params to pass in format of name=value. Can specify multiple times")
	cmd.Flags().StringVarP(&contentType, "content-type", "c", "application/json", "content-type to return")
	rootCmd := cli.New(cmd)
	_ = rootCmd.Execute()
}
