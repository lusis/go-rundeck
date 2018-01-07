# kingpin examples

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

```text
+---------+---------+---------+
| HEADER1 | HEADER2 | HEADER3 |
+---------+---------+---------+
| value1  | value2  | value3  |
+---------+---------+---------+
```