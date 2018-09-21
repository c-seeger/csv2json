package main

import (
	"flag"
	"log"
	"strings"

	"github.com/cseeger-epages/csv2json"
)

func main() {
	path := flag.String("c", "./csv/somename_19700101.csv", "path to csv file")
	output := flag.String("o", "./json/somename_19700101.json", "path to output file")
	flag.Parse()

	// parse info from filename and add them to json
	addFields := make(map[string]string)

	// we want the name and the date from our file
	data := strings.Split(*path, "/")
	fields := strings.Split(data[len(data)-1], "_")

	addFields["name"] = fields[0]
	addFields["date"] = fields[1][:len(fields[1])-4] // len(fields[1])-4 removes the .csv from the string
	addFields["filepath"] = *path

	// adding some fields
	fh, err := csv2json.ReadFile(path)
	if err != nil {
		log.Fatal(err)
	}
	fileBytes, err := csv2json.ReadCSV(fh, addFields, csv2json.Options{})
	if err != nil {
		log.Fatal(err)
	}
	err = csv2json.SaveFile(fileBytes, *output)
	if err != nil {
		log.Fatal(err)
	}
}
