# go-rundeck
[![Build Status](https://travis-ci.org/lusis/go-rundeck.svg?branch=master)](https://travis-ci.org/lusis/go-rundeck)
Go library and utilities for interacting with [Rundeck](http://rundeck.org).
It currently supports configuration via environment variables only and supports either token authentication or username/password authentication. (v13 only)

# Configuration
- set `RUNDECK_URL` to the base url of your rundeck installation (i.e. `http://localhost:4000` or `https://my.rundeck.domain.com`)

- set *EITHER* `RUNDECK_TOKEN` or `RUNDECK_USERNAME` and `RUNDECK_PASSWORD` (only if using `v13` api)
- if all three are set `RUNDECK_TOKEN` takes precendence

Again, username/password authentication are ONLY supported in `rundeck.v13` and higher. The original `rundeck.v12` only supports token authentication.

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

	rundeck "github.com/lusis/go-rundeck/src/rundeck.v17"
)

func main() {
	/*
		NewClientFromEnv requires two (or three) environement variables:
		- RUNDECK_URL
		- RUNDECK_TOKEN
		if RUNDECK_TOKEN is empty, RUNDECK_USERNAME/RUNDECK_PASSWORD are required
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

## Note on versioning
The initial release of this repo was built for rundeck API v12. Each rundeck release adds new or changed functionality.
I'm now using a v17 install of Rundeck. The v13 API is the last before several changes in the v14 api.

For now I plan on keeping a different directory for EACH version on the master branch even though there will be quite a bit of code duplication. I reserve the right to make changes to the internals of the library as well in each versioned directory. Some of this is happening with v13 as I ended up shooting myself in the foot with the initial implementation and didn't leave room to change with the rundeck API.

The question becomes how to handle the bundled utilities. I'm likely going to break those out into separate repos and version them separately.

## TODO
- Flesh out more tests
- More utilities
- Wrapper cli for sub-utilities
