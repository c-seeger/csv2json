/*
Package csv2json implementing a simple csv2 json converter

Simple example:

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
*/
package csv2json
