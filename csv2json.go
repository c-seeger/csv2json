package csv2json

import (
	"encoding/csv"
	"fmt"
	"os"
)

// Options to controll the json generation
type Options struct {
	LineWiseJSON    bool
	PrettyPrint     bool
	QuoteEverything bool
}

// ReadCSV to read the content of CSV File
func ReadCSV(csvFile *os.File, additionalFields map[string]string, opt Options) ([]byte, error) {
	defer csvFile.Close()

	reader := csv.NewReader(csvFile)
	content, err := reader.ReadAll()
	if err != nil {
		return nil, err
	}

	if len(content) < 1 {
		return nil, fmt.Errorf("Something wrong, the file maybe empty or length of the lines are not the same")
	}

	if opt.LineWiseJSON {
		return lineWiseJSON(content, additionalFields, opt)
	}
	return intoJSONArray(content, additionalFields, opt)
}
