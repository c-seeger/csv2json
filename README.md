# csv2json

is a [CSV](https://en.wikipedia.org/wiki/Comma-separated_values) to [JSON](https://en.wikipedia.org/wiki/JSON) converter inspired by the work of [CSV-To-JSON-Converter](https://github.com/Ahmad-Magdy/CSV-To-JSON-Converter).

The original project was not a library implementation so i encapsulated the code into the csv2json package and added some features to create a more usable version.

## Features
- convert csv to json
- adding additional fields that are not in the original csv file

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
        fileBytes, err := csv2json.ReadCSV(path, nil)
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

### Advanced examples

see [examples](https://github.com/cseeger-epages/csv2json/tree/master/examples) for some more usage examples
