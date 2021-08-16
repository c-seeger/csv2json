# csv2json

[![Donate](https://img.shields.io/badge/Donate-PayPal-green.svg)](https://www.paypal.com/cgi-bin/webscr?cmd=_s-xclick&hosted_button_id=VBXHBYFU44T5W&source=url)
[![GoDoc](https://img.shields.io/badge/godoc-reference-green.svg)](https://godoc.org/github.com/c-seeger/csv2json)
[![Go Report Card](https://goreportcard.com/badge/github.com/c-seeger/csv2json)](https://goreportcard.com/report/github.com/c-seeger/csv2json)
[![License](https://img.shields.io/badge/license-MIT-blue.svg)](https://github.com/c-seeger/csv2json/blob/master/LICENSE)
[![Build Status](https://travis-ci.org/c-seeger/csv2json.svg?branch=master)](https://travis-ci.org/c-seeger/csv2json)
[![Built with Mage](https://magefile.org/badge.svg)](https://magefile.org)

is a [CSV](https://en.wikipedia.org/wiki/Comma-separated_values) to [JSON](https://en.wikipedia.org/wiki/JSON) converter inspired by the work of [CSV-To-JSON-Converter](https://github.com/Ahmad-Magdy/CSV-To-JSON-Converter).

This project was created by the intension to create json files from csv for importing the data into an ELK stack since JSON supports more options than CSV for ELK.

## Features

- convert csv to json
- adding additional fields that are not in the original csv file
- options PrettyPrint, QuoteEverything and LineWiseJSON

## Installation

If you already installed GO on your system and configured it properly than its simply:

```
go get github.com/c-seeger/csv2json
```

If not follow [these instructions](https://nats.io/documentation/tutorials/go-install/).

## Usage

### Simple example

The following example reads a csv file defined by -c parameter and outputs to a file defined by -o

```
func main() {
  // read and convert csv
  f, err := os.Open("test/fixtures/data.csv")
  if err != nil {
    log.Fatal(err)
  }
  defer f.Close()

  // optional add additional fields
  addFields := make(map[string]string)
  addFields["addedField"] = "some data"

  bytes, err := csv2json.ReadCSV(f, addFields, csv2json.Options{
    PrettyPrint:     true,
    LineWiseJSON:    false,
    QuoteEverything: false,
  })
  if err != nil {
    log.Fatal(err)
  }

  fmt.Println(string(bytes))
}

```

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

## Contribution

Thank you for participating to this project.
Please see our [Contribution Guidlines](https://github.com/c-seeger/csv2json/blob/master/CONTRIBUTING.md) for more information.
