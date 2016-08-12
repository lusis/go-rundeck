package main

import (
	"flag"
	"fmt"
	"os"
	"strings"

	rundeck "rundeck.v17"
)

var usage = `
Returns xml dump of rundeck api calls. Useful for dumping xml for test cases.

Options:
	-path="projects": path to dump (i.e. executions/12234)
	-q=[]: key=value query parameter. Specify multiple times if neccessary

Usage: rundeck-xml-get [-q queryopt=queryvalue] -path [path]
	
	List projects:
		rundeck-xml-get -path projects
	List jobs for project:
		rundeck-xml-get -q project=MYPROJ -path jobs
	List at most two jobs:
		rundeck-xml-get -q project=MYPROJ -q max=2 -path jobs
`

type paramsType []string

func (p *paramsType) Set(value string) error {
	*p = append(*p, value)
	return nil
}

func (p *paramsType) String() string {
	return fmt.Sprintf("%s", *p)
}

func paramConvert(p paramsType, pmap *map[string]string) {
	if len(p) != 0 {
		localmap := *pmap
		for _, param := range p {
			arr := strings.Split(param, "=")
			localmap[arr[0]] = arr[1]
		}
	}
}

var myparams paramsType

func main() {
	path := flag.String("path", "projects", "path to dump (i.e. executions/12234).")
	flag.Var(&myparams, "q", "key=value query parameter. Specify multiple times if neccessary")
	flag.Usage = func() { fmt.Printf("%s", usage) }
	flag.Parse()
	query_opts := make(map[string]string)
	paramConvert(myparams, &query_opts)
	client := rundeck.NewClientFromEnv()
	var data []byte
	err := client.Get(&data, *path, query_opts)
	if err != nil {
		fmt.Printf(err.Error())
		os.Exit(1)
	} else {
		fmt.Printf("%s\n", string(data))
	}
}
