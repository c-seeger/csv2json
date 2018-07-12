# csv2json

is a [CSV](https://en.wikipedia.org/wiki/Comma-separated_values) to [JSON](https://en.wikipedia.org/wiki/JSON) converter inspired by the work of [CSV-To-JSON-Converter](https://github.com/Ahmad-Magdy/CSV-To-JSON-Converter).

The original project was not a library implementation so i encapsulated the code into the csv2json package and added some features to create a more usable version.

This project was created by the intension to create json files from csv for importing the data into an ELK stack since JSON supports more options than CSV for ELK.

## Features
- convert csv to json
- adding additional fields that are not in the original csv file
- options PrettyPrint, QuoteEverything and LineWiseJSON

## Installation

If you already installed GO on your system and configured it properly than its simply:

```
go get github.com/cseeger-epages/csv2json
```

If not follow [these instructions](https://nats.io/documentation/tutorials/go-install/).

## Usage

### Simple example

The following example reads a csv file defined by -c parameter and outputs to a file defined by -o 

```
//test.go
package main

import (
        "flag"
        "log"

        "github.com/cseeger-epages/csv2json"
)

func main() {
        path := flag.String("c", "./data.csv", "path of the file")
        output := flag.String("o", "./data.json", "path of the output file")
        flag.Parse()
        fileBytes, err := csv2json.ReadCSV(path, nil, csv2json.Options{})
        if err != nil {
                log.Fatal(err)
        }
        err = csv2json.SaveFile(fileBytes, *output)
        if err != nil {
                log.Fatal(err)
        }
}
```

run it by `go run test.go -c test.csv -o test.json` or use `go build` to create a binary

### Options

There are some options available in the `csv2json.Options` struct:

```
type Options struct {
        LineWiseJSON    bool
        PrettyPrint     bool
        QuoteEverything bool
}
```

- LineWiseJSON
  - creates a json line per csv line and concats all lines together in one file
  - this feature is usefull if you want to generate JSON files from csv's as a data source for an ELK stack
  - if this is set to false a jSON array with every line as an entry will be created 
- PrettyPrint
  - more readable multiline version of json
- QuoteEverything
  - ignores any kind of datatypes and quotes everything

### Advanced examples

see [examples](https://github.com/cseeger-epages/csv2json/tree/master/examples) for some more usage examples
