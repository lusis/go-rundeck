# outputter

[![Build Status](https://travis-ci.org/lusis/outputter.svg?branch=master)](https://travis-ci.org/lusis/outputter)

Outputter is a small go library for formatting data you want to present to stdout.
The feel is very much modeled after [tablewriter](github.com/olekukonko/tablewriter)

## usage

In the `example` directory, there are examples for `flag`, `kingpin` and `cobra` usages.
If your arg parsing library let's you specify a list of valid options for an argument, you can call `outputter.GetOutputters()` to get a `[]string` of all the registered outputters.
Ideally a user should be able to pass whatever format is specified on the command line directly into `outputter.NewOutputter()` as in the `kingpin` example:

```go
    kingpin.Flag("format", "format for output").
        Default("tabular").
        EnumVar(&outputFormat, outputter.GetOutputters()...)
    kingpin.Parse()
    outputFormatter, err := outputter.NewOutputter(outputFormat)
```

If an invalid format is specified, `ErrorInvalidOutputter` is returned.

## color support

If you like to use the awesome [color](https://github.com/fatih/color) package, you can but it's up to each output formatter to decide if it supports colorized output or not. The best way to handle this is like so:

```go
color.NoColor = !outputFormatter.ColorSupport()
outputFormatter.SetHeaders([]string{
    "header1",
    "header2",
    "header3",
})
rowErr := outputFormatter.AddRow([]string{
    color.YellowString("value1"),
    color.YellowString("value2"),
    color.YellowString("value3"),
})
```

It's important to trigger the flag BEFORE you start setting color codes.

There's an example in `example/color`

## pretty output

If an output format supports a pretty print mechanism of some kind, you can call `outputFormatter.SetPretty()` to trigger it.

Currently only the `json` output supports that.

## note about json output

Depending on your shell configuration and prompts, a single line json may get "erased" when outputting. You see this frequently with complicated zsh prompts. You can be sure your output is working by piping through `jq`.

## provided output formats

### tabular

Draws results using `tabwriter`:

```text
header1  header2  header3
value1   value2   value3
```

_Does not support colorized output due to a bug in tabwriter_

### table

Uses [tablewriter](github.com/olekukonko/tablewriter) to draw results as a table

```text
+---------+---------+---------+
| HEADER1 | HEADER2 | HEADER3 |
+---------+---------+---------+
| value1  | value2  | value3  |
+---------+---------+---------+
```

_Supports colorized output_

### json

Writes results as json representing each row as an object in an array with headers as key names

```json
[{"header1":"value1","header2":"value2","header3":"value3"}]
```

Supports pretty printed output like so:

```json
[
    {
        "header1": "value1",
        "header2": "value2"
    },
    {
        "header1": "value3",
        "header2": "value4"
    }
]
```

_does not support colorized output_

## jsonshort

A newer json output that might make more sense.

```json
{"header1": ["value1","value3"],"header2": ["value2","value4"]}
```

Supports a prettyprint version like so:

```json
{
    "key1": [
        "value1",
        "value3"
    ],
    "key2": [
        "value2",
        "value4"
    ]
}
```

_does not support colorized output_

## extending

The requirements to be an outputter are pretty minimal.

In your package you'll need to call `outputter.RegisterOutput("output-name", customOutputFactory)`
You can get an instance of it to use with `outputter.NewOutputter("output-name")`

Ideally you should support both a `stdout` and an `io.Writer` version of your output to make testing feasible. The default should be the `stdout` version.

You should create whatever actually uses your `io.Writer` as late as possible (ideally in `Draw()`) to ensure that users can call `SetWriter()` at any point up until `Draw()` is called.
Ideally, users should even be able to call `Draw()`, change the writer and call `Draw()` again to another writer.
In cases where this is simply not possible, you should return `ErrorCannotChangeWriter` in your `SetWriter()` function based on whatever criteria triggers that (i.e. headers already set, rows already populated)

There are some error constants you should use defined in `errors.go` and return.

You can see an example in `example/custom-ouput`