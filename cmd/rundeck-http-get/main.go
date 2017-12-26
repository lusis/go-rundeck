package main

import (
	"fmt"
	"log"
	"os"
	"strings"

	httpclient "github.com/lusis/go-rundeck/pkg/httpclient"
	rundeck "github.com/lusis/go-rundeck/pkg/rundeck.v21"
	kingpin "gopkg.in/alecthomas/kingpin.v2"
)

var (
	path            = kingpin.Arg("path", "path to dump (e.g. executions/12234)").Required().String()
	queryParameters = kingpin.Flag("query_params", "key=value query parameter. specify multiple times if neccessary").Strings()
	contentType     = kingpin.Flag("content_type", "an alternate content type if neccessary").Default("application/json").String()
)

func buildParams(p *map[string]string, value string) error {
	parts := strings.SplitN(value, "=", 2)
	if len(parts) != 2 {
		return fmt.Errorf("expected key=value got '%s'", value)
	}
	(*p)[parts[0]] = parts[1]
	return nil
}

func main() {
	myParams := make(map[string]string)
	kingpin.Parse()
	client, clientErr := rundeck.NewClientFromEnv()

	if clientErr != nil {
		log.Fatal(clientErr.Error())
	}
	options := []httpclient.RequestOption{}
	if contentType != nil {
		options = append(options, httpclient.Accept(*contentType))
	}

	for _, param := range *queryParameters {
		e := buildParams(&myParams, param)
		if e != nil {
			fmt.Printf(e.Error())
			os.Exit(1)
		}
		options = append(options, httpclient.QueryParams(myParams))
	}
	data, err := client.Get(*path, options...)
	if err != nil {
		fmt.Printf(err.Error())
		os.Exit(1)
	} else {
		fmt.Printf("%s\n", string(data))
	}
}
