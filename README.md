# go-rundeck

[![Build Status](https://img.shields.io/circleci/project/github/lusis/go-rundeck/master.svg?style=for-the-badge)](https://circleci.com/gh/lusis/go-rundeck)

Go library and utilities for interacting with [Rundeck](http://rundeck.org).

## ** Potentially breaking changes coming soon **

If you are using this library without any sort of vendoring or dependency management, you need to start now.

As of the v21 changes, I've started down a path that will allow me to handle api changes better:

Starting with v21, there is no longer a versioned import path. It is now simply `github.com/lusis/go-rundeck/pkg/rundeck`
As of the next release of the API, I will be removing ALL of the non-current rundeck API versioned import paths from `master`.
This means if you are NOT vendoring and pinning to a specific SHA, your stuff will break.

I've not finalized how I will be handling API changes but maintaining all the old versions in a single master has become very difficult.

### Isolation request and response from golang api

I've separated requests and responses from the actual golang rundeck API. Those will potentially the only thing versioned going forward.

There was a massive tangled mess where the rundeck golang types were too tightly coupled to the actual request and response from the Rundeck API.
By breaking these out, it's really made things much more flexible AND testable.

### Actual version support

I've added min/max version support around funcs that perform API operations that can help with a more unified codebase.
I tried very hard to specify the absolute minimum version a user could get away with but in some cases where there were breaking api changes,
I chose to fight the battle another day and specified a bit higher minimum version constraint.

I've also added environment variable support for specifying the version of the rundeck API you want to use. This can be used with the new unified `rundeck` binary

### unified rundeck binary

One of the pain points I faced as a maintainer was updating all the individual binaries when I migrated to a new API release.
Combined with the new version constraints in functions and the `RUNDECK_VERSION` environment variable, I can now offer a binary that should hopefully work across versions of the rundeck server.

As an example:

```text
$ rundeck logstorage get
+----------+-------------+-----------+--------+--------+-------+------------+---------+
| ENABLED? | PLUGIN NAME | SUCCEEDED | FAILED | QUEUED | TOTAL | INCOMPLETE | MISSING |
+----------+-------------+-----------+--------+--------+-------+------------+---------+
| false    |             | 0         | 0      | 0      | 0     | 0          | 1       |
+----------+-------------+-----------+--------+--------+-------+------------+---------+

$ RUNDECK_VERSION=14 rundeck logstorage get
Error: Requested API version (14) does not meet the requirements for this api call (min: 17, max: 21)
```

## Configuration

- set `RUNDECK_URL` to the base url of your rundeck installation (i.e. `http://localhost:4000` or `https://my.rundeck.domain.com`)
- set *EITHER* `RUNDECK_TOKEN` or `RUNDECK_USERNAME` and `RUNDECK_PASSWORD` (only if using `v13` api)
- if all three are set `RUNDECK_TOKEN` takes precendence
- `RUNDECK_VERSION` can be used if you're running a lower version of the rundeck server api but nothing has changed in newer versions.

## Usage

There are two ways to use this:

- as a library
- request/response types
- via the unified binary, `rundeck`

### as a library

```go
package main

import (
    "fmt"
    rundeck "github.com/lusis/go-rundeck/pkg/rundeck"
)

func main() {
    client, clientErr := rundeck.NewClientFromEnv()
    if clientErr != nil {
        log.Fatal(clientErr)
    }
    data, err := client.ListProjects()
    if err != nil {
        fmt.Println(err)
    } else {
        fmt.Printf("%+v\n", data)
    }
}
```

### request/response types

As part of some design changes to the library, you can now just import the request/response types.
If you don't like my golang API and want to write your own, you can still benefit from the work I did ;)

```go
package main

import (
    "encoding/json"

    "github.com/davecgh/go-spew/spew"
    responses "github.com/lusis/go-rundeck/pkg/rundeck.v21/responses"
)

// SystemInfo represents the rundeck server system info output
type SystemInfo responses.SystemInfoResponse

// GetSystemInfo gets system information from the rundeck server
// http://rundeck.org/docs/api/index.html#system-info
func main() {
    ls := SystemInfo{}
    // data is populated via some mechanism where you get rundeck api json
    jsonErr := json.Unmarshal(data, &ls)
    if jsonErr != nil {
        log.Fatal(jsonErr)
    }
    spew.Dump(ls)
}
```

### `rundeck` unified binary

```text
$ go get github.com/lusis/go-rundeck/cmd/rundeck
$ RUNDECK_URL=http://rundeck.local:4440 RUNDECK_TOKEN=XXXXXXX rundeck help
Unified rundeck cli binary

Usage:
  rundeck [command]

Available Commands:
  execution   operate on individual rundeck executions
  executions  operate on rundeck multiple rundeck executions at once
  help        Help about any command
  http        perform authenticated http operations against a rundeck server. kinda like curl
  job         operate on individual rundeck jobs
  jobs        operate on rundeck multiple rundeck jobs at once
  list        list various things from the rundeck server
  logstorage  operate on rundeck logstorage
  policy      operate on rundeck acl policies
  project     operate on a rundeck project
  token       operate on an individual token in rundeck
  tokens      operate on rundeck api tokens
```
