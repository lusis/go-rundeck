package main

import (
	"fmt"
	"log"
	"os"

	rundeck "github.com/lusis/go-rundeck/pkg/rundeck.v21"
	kingpin "gopkg.in/alecthomas/kingpin.v2"
)

var (
	token = kingpin.Arg("token", "token to delete").Required().String()
)

func main() {
	kingpin.Parse()
	client, clientErr := rundeck.NewClientFromEnv()
	if clientErr != nil {
		log.Fatal(clientErr.Error())
	}
	err := client.DeleteToken(*token)
	if err != nil {
		fmt.Printf("%s\n", err)
		os.Exit(1)
	} else {
		fmt.Printf("Token deleted\n")
		os.Exit(0)
	}
}
