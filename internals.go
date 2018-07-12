package csv2json

import (
	"bytes"
	"encoding/json"
	"strconv"
	"strings"
)

func intoJSONArray(content [][]string, additionalFields map[string]string, opt Options) ([]byte, error) {
	headersArr := getHeader(content)
	//Remove the header row
	content = content[1:]

	var (
		buffer bytes.Buffer
	)

	buffer.WriteString("[")
	for i, d := range content {
		buffer.WriteString("{")
		for j, y := range d {
			buffer.WriteString(`"` + headersArr[j] + `":`)
			if opt.QuoteEverything {
				buffer.WriteString((`"` + y + `"`))
			} else {
				_, fErr := strconv.ParseFloat(y, 32)
				_, bErr := strconv.ParseBool(y)
				if fErr == nil {
					buffer.WriteString(y)
				} else if bErr == nil {
					buffer.WriteString(strings.ToLower(y))
				} else {
					buffer.WriteString((`"` + y + `"`))
				}
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
	if opt.PrettyPrint {
		return json.MarshalIndent(rawMessage, "", "  ")
	}
	return json.Marshal(rawMessage)
}

func lineWiseJSON(content [][]string, additionalFields map[string]string, opt Options) ([]byte, error) {
	headersArr := getHeader(content)
	//Remove the header row
	content = content[1:]

	var retJSON []byte
	nl := []byte("\n")

	for _, d := range content {
		var (
			buffer bytes.Buffer
			js     []byte
			err    error
		)
		buffer.WriteString("{")
		for j, y := range d {
			buffer.WriteString(`"` + headersArr[j] + `":`)
			if opt.QuoteEverything {
				buffer.WriteString((`"` + y + `"`))
			} else {
				_, fErr := strconv.ParseFloat(y, 32)
				_, bErr := strconv.ParseBool(y)
				if fErr == nil {
					buffer.WriteString(y)
				} else if bErr == nil {
					buffer.WriteString(strings.ToLower(y))
				} else {
					buffer.WriteString((`"` + y + `"`))
				}
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
		rawMessage := json.RawMessage(buffer.String())
		if opt.PrettyPrint {
			js, err = json.MarshalIndent(rawMessage, "", "  ")
			if err != nil {
				return nil, err
			}
		} else {
			js, err = json.Marshal(rawMessage)
			if err != nil {
				return nil, err
			}
		}
		js = append(js, nl...)

		retJSON = append(retJSON, js...)
	}

	return retJSON, nil

}

func getHeader(content [][]string) []string {
	headersArr := make([]string, 0)
	for _, headE := range content[0] {
		headersArr = append(headersArr, headE)
	}

	return headersArr
}
