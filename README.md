# go-rundeck
Go library and utilities for interacting with [Rundeck](http://rundeck.org)

## Usage
There are two ways to use this:
- as a library
- via the bundled utilities

### as a library
```go
package main

import (
	"fmt"

	rundeck "github.com/lusis/go-rundeck/src/rundeck.v12"
)

func main() {
	/*
		NewClientFromEnv requires two environement variables:
		- RUNDECK_URL
		- RUNDECK_TOKEN
	*/
	client := rundeck.NewClientFromEnv()
	data, err := client.ListProjects()
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("%+v\n", data)
	}
}
```

### bundled utilities
```
git clone https://github.com/lusis/go-rundeck.git
cd go-rundeck
make all
RUNDECK_URL=http://rundeck.local:4440 RUNDECK_TOKEN=XXXXXXX bin/rundeck-list-projects
```

```
+---------+-------------+-----------------------------------------------------------+
|  NAME   | DESCRIPTION |                            URL                            |
+---------+-------------+-----------------------------------------------------------+
| FOOPROJ | A project   | http://rundeck.local:4440/api/12/project/FOOPROJ          |
+---------+-------------+-----------------------------------------------------------+
```
