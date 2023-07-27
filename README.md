# goday - A Go package for Date Manipulation and Formatting

[![Go Report Card](https://goreportcard.com/badge/github.com/deryl-sagala/goday)](https://goreportcard.com/report/github.com/deryl-sagala/goday)[![Go Reference](https://pkg.go.dev/badge/pkg.go.dev/github.com/deryl-sagala/goday.svg)](https://pkg.go.dev/pkg.go.dev/github.com/deryl-sagala/goday)

The `goday` package is a Date and Time manipulation library for Go (Golang) inspired by Day.js, a popular JavaScript library for date manipulation. This package provides functionalities to parse, format, manipulate, and query dates in a simple and intuitive way, similar to Day.js.

## Installation

```bash
go get github.com/deryl-sagala/goday
```

## Usage

Here's an overview of the main functionalities provided by the `goday` package:

### Parsing Dates

You can parse a date string into a `goday` instance using the `goday.Parse()` function.

```go
import "github.com/deryl-sagala/goday"

date, err := goday.Parse("2018-08-08")
if err != nil {
    // Handle parsing error
}
```

### Formatting Dates

Use the `goday.Format()` function to format dates into custom layouts.

```go
formattedDate := date.Format("{YYYY}-{MM}-{DD} {HH:mm:ss}")
```

### Get & Set Date Components

The `goday.Get()` and `goday.Set()` functions allow you to retrieve or set specific date components.

```go
year := date.Get("year")
date.Set("month", 3) // Set the month to April (3)
```

### Date Manipulation

Easily manipulate dates by adding or subtracting time units.

```go
oneYearLater := date.Add(1, "year")
threeDaysAgo := date.Subtract(3, "day")
```

### Date Querying

Use `goday.IsBefore()`, `goday.IsAfter()`, and other functions to query dates.

```go
isBefore := date1.IsBefore(date2)
isAfter := date1.IsAfter(date2)
isSame := date1.IsSame(date2, "year")
```

## Contributions

Contributions are welcome! If you find any issues or have suggestions for improvements, feel free to create an issue or submit a pull request.

## License

This package is distributed under the Apache 2.0 License. See the [LICENSE](LICENSE) file for more information.

Happy Date Manipulation with `goday` in Go!
