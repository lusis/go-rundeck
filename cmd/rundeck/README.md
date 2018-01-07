# `rundeck` unified cli binary

As of `v21` of this library, we've migrated to a single binary replacing all the previous individual binaries.
This results in much faster builds and is frankly easier to manage common logic in one place.

## Install

Heyo it's go gettable! (or it will be when this branch is merged down to master =P)

## Usage

The command really does walk you through all the options but in general the flow is like so:

- subcommands for "things" in rundeck in singular and plural forms as makes sense (i.e. `rundeck job run` vs `rundeck jobs` for bulk operations)
- a few shortcut aliases like `rundeck list jobs <projectname>` vs `rundeck project jobs <projectname>` to be a bit more ergonomic
- CRUD verbs where it makes sense (i.e. `rundeck job get`/`rundeck job create`/`rundeck job delete`)
- minimum required options are usually the bare arguments (i.e. `rundeck job get <jobid>`) vs flags
- universal output formatting via [outputter](https://github.com/lusis/outputter) (under the `--output-format` flag) except where explicitly disabled (i.e. exporting a job definition or anywhere the raw body of the response is what we want)

## TODO

While all the previous commands exist in the unified binary now, one feature in I've not yet ported is the ability to just dump ids from commands that can be used to pipe between themselves. This may not be as critical with the various json outputs and jq but it's a thing I'd like to have

I also have quite a bit more API functionlity that we now have exposed in the library and bring it into the cli.

I also need to finish the other `http` subcommands.

## Sample help output

```text
$ rundeck help
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

Flags:
  -h, --help   help for rundeck

Use "rundeck [command] --help" for more information about a command.
```

```text
$ rundeck list -h
list various things from the rundeck server

Usage:
  rundeck list [command]

Available Commands:
  executions  gets a list of executions for a project from the rundeck server optionally only running executions
  jobs        lists all jobs for a project
  projects    gets a list of projects from the rundeck server

Flags:
  -h, --help   help for list

Use "rundeck list [command] --help" for more information about a command
```

```text
$ rundeck project get -h
rundeck project get -h
gets project info from the rundeck server

Usage:
  rundeck project get project-name [flags]

Flags:
  -h, --help                   help for get
      --output-format string   Specify the output format: csv,json,jsonshort,table,tabular (default "table")
```

```text
$ rundeck http get -h
rundeck http get -h

"rundeck http get" allows you to perform authenticated http get requests against the rundeck api.
Examples:

# rundeck http get system/acl/
# rundeck http get job/XXXXXXX/executions -q max=1
# rundeck http get job/XXXXXXX/executions -q max=1 -q status=failed
# rundeck http get execution/29 -c application/xml

This tool is used to generate test response data for this library itself.

Usage:
  rundeck http get path [-q foo=bar] [-c application/json] [flags]

Flags:
  -c, --content-type string       content-type to return (default "application/json")
  -h, --help                      help for get
  -q, --query-param stringSlice   custom query params to pass in format of name=value. Can specify multiple times
```

## Migration to `cobra`

The binary now uses [cobra](https://github.com/spf13/cobra) instead of kingping.
Obviously migrating all of this was no easy chore so you might be wondering why.

The first reason was mainly consistency with my day to day work.
We settled on cobra internally at Mailchimp and I really didn't like the context switching I was doing.

The second reason is traction. A lot of popular libraries and utilities are using cobra and so the flow is somewhat well documented across projects.

I want to say that I'm still a fan of [kingpin](https://github.com/alecthomas/kingpin).

I guess if you wanted a quicklist of my thoughts:

### kingpin

- better documentation for complex usage and customization
- better out-of-box argument types such as `Enum` for enforcing options next to the option definition
- [personal opinion] less ergonomic api due to function chaining/fluent style

### cobra

- building critical mass
- I really like the "`Run`/`RunE`/`PreRunE`/`PreRun`/`PersistantPreRun`/you get the idea" options (though I am over huge structs these days)
- [personal opinion] Friendlier API
- too much `init()` magic as default examples
- less flexible default types in `spf13/pflag` library it uses
- worse documentation for complex usage and customization (i.e. I hate `init()` magic)
- lack of a `map[string]string` type is frustrating
