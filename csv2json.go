package csv2json

import (
	"bytes"
	"encoding/csv"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
)

// ReadCSV to read the content of CSV File
func ReadCSV(path *string, additionalFields map[string]string) ([]byte, error) {
	csvFile, err := os.Open(*path)
	if err != nil {
		return nil, err
	}
	defer csvFile.Close()

	reader := csv.NewReader(csvFile)
	content, _ := reader.ReadAll()

	if len(content) < 1 {
		return nil, fmt.Errorf("Something wrong, the file maybe empty or length of the lines are not the same")
	}

	headersArr := make([]string, 0)
	for _, headE := range content[0] {
		headersArr = append(headersArr, headE)
	}

	//Remove the header row
	content = content[1:]

	var buffer bytes.Buffer
	buffer.WriteString("[")
	for i, d := range content {
		buffer.WriteString("{")
		for j, y := range d {
			buffer.WriteString(`"` + headersArr[j] + `":`)
			_, fErr := strconv.ParseFloat(y, 32)
			_, bErr := strconv.ParseBool(y)
			if fErr == nil {
				buffer.WriteString(y)
			} else if bErr == nil {
				buffer.WriteString(strings.ToLower(y))
			} else {
				buffer.WriteString((`"` + y + `"`))
			}
			//end of property
			if j < len(d)-1 {
				buffer.WriteString(",")
			} else if len(additionalFields) > 0 {
				buffer.WriteString(",")
			}

		}
		ai := 0
		for k, v := range additionalFields {
			buffer.WriteString(`"` + k + `":"` + v + `"`)
			if ai < len(additionalFields)-1 {
				buffer.WriteString(",")
			}
			ai++
		}
		//end of object of the array
		buffer.WriteString("}")
		if i < len(content)-1 {
			buffer.WriteString(",")
		}
	}

	buffer.WriteString(`]`)
	rawMessage := json.RawMessage(buffer.String())
	json, err := json.MarshalIndent(rawMessage, "", "  ")
	if err != nil {
		return nil, err
	}
	return json, nil
}

// SaveFile creates a file by a given content and path
func SaveFile(content []byte, path string) error {
	if err := ioutil.WriteFile(path, content, os.FileMode(0644)); err != nil {
		return err
	}
	return nil
}
