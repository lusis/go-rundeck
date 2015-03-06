# go-rundeck
Go library and utilities for interacting with [Rundeck](http://rundeck.org)

## Usage
There are two ways to use this:
- as a library
- just for the unmarshal
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

### for unmarshalling existing data into structs
_This is what the test cases do with the mock data_

```go
package main

import (
        "encoding/xml"
        "io/ioutil"
        "os"
)

	xmlfile, err := os.Open("assets/test/user_token.xml")
	if err != nil {
		t.Fatal(err.Error())
	}
	defer xmlfile.Close()
	xmlData, _ := ioutil.ReadAll(xmlfile)
	var s Tokens
	xml.Unmarshal(xmlData, &s)

	fmt.Printf("%+v", s)
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

## TODO
- Flesh out more tests
- Add mock for http
- More utilities
- Wrapper cli for sub-utilities
