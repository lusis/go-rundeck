# colorized examples
In general, you should be able to pass a color code to the output and things will print properly.

However color makes no sense for json and there's a bug for tabwriter preventing it from working

This example shows how to use `fatih/color`'s `NoColor` in conjunction.

## tabular

`go run main.go`

```text
header1  header2  header3
value1   value2   value3
```

## json

`go run main.go --format json`

```text
[{"header1":"value1","header2":"value2","header3":"value3"}]
```

## table

`go run main.go --format table`

![COLORZ!](color-table.png?raw=true)

```text

```