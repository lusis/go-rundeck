package main

import (
	"fmt"
	"os"

	rundeck "github.com/lusis/go-rundeck/pkg/rundeck.v19"
	kingpin "gopkg.in/alecthomas/kingpin.v2"
)

var (
	user = kingpin.Arg("username", "Username to assign the new token").Required().String()
)

func main() {
	kingpin.Parse()
	client := rundeck.NewClientFromEnv()

	token, err := client.CreateToken(*user)
	if err != nil {
		fmt.Printf("%s\n", err)
		os.Exit(1)
	} else {
		fmt.Printf("%s\n", token)
		os.Exit(0)
	}
}
