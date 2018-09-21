package main

import (
	"flag"
	"log"

	"github.com/cseeger-epages/csv2json"
)

func main() {
	// define csv path and output path
	csvPath := flag.String("c", "./csv/data.csv", "path to csv file")
	outputPath := flag.String("o", "./json/data.json", "path to output file")
	flag.Parse()

	// read and convert csv
	fh, err := csv2json.ReadFile(csvPath)
	if err != nil {
		log.Fatal(err)
	}
	fileBytes, err := csv2json.ReadCSV(fh, nil, csv2json.Options{})
	if err != nil {
		log.Fatal(err)
	}
	err = csv2json.SaveFile(fileBytes, *outputPath)
	if err != nil {
		log.Fatal(err)
	}
}
