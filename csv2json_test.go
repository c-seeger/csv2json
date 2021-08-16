package csv2json

import (
	"encoding/json"
	"io/ioutil"
	"os"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestReadCSVNoFileHandle(t *testing.T) {
	assert := assert.New(t)

	_, err := ReadCSV(nil, nil, Options{})
	assert.Error(err)
}

func TestReadCSVEmptyFileHandle(t *testing.T) {
	assert := assert.New(t)

	file, err := ioutil.TempFile("", "empty-file")
	assert.NoError(err)
	defer file.Close()
	defer os.Remove(file.Name())

	_, err = ReadCSV(file, nil, Options{})
	assert.Error(err)
}

func TestReadCSV(t *testing.T) {
	assert := assert.New(t)

	file, err := os.Open("test/fixtures/data.csv")
	assert.NoError(err)

	res, err := ReadCSV(file, nil, Options{})
	assert.NoError(err)

	json, err := ioutil.ReadFile("test/fixtures/data.json")
	assert.NoError(err)

	assert.Equal(json, res)
}

func TestReadCSVAdditionalFields(t *testing.T) {
	assert := assert.New(t)

	file, err := os.Open("test/fixtures/data.csv")
	assert.NoError(err)

	addFields := make(map[string]string)
	addFields["addedField"] = "test"

	res, err := ReadCSV(file, addFields, Options{})
	assert.NoError(err)

	var jsmap []map[string]interface{}

	err = json.Unmarshal(res, &jsmap)
	assert.NoError(err)

	_, ok := jsmap[0]["addedField"]
	assert.True(ok)
	if ok {
		assert.Equal(jsmap[0]["addedField"], "test")
	}
}

func TestReadCSVOptionsLineWiseJSON(t *testing.T) {
	assert := assert.New(t)

	file, err := os.Open("test/fixtures/data.csv")
	assert.NoError(err)

	res, err := ReadCSV(file, nil, Options{LineWiseJSON: true})
	assert.NoError(err)

	json, err := ioutil.ReadFile("test/fixtures/linewise.json")
	assert.NoError(err)

	assert.Equal(json, res)
}

func TestReadCSVOptionsQuoteEverything(t *testing.T) {
	assert := assert.New(t)

	file, err := os.Open("test/fixtures/data.csv")
	assert.NoError(err)

	res, err := ReadCSV(file, nil, Options{QuoteEverything: true})
	assert.NoError(err)

	// if everything is quoted we can map it to string
	// instead of specific types without creating errors
	var jsmap []map[string]string
	err = json.Unmarshal(res, &jsmap)
	assert.NoError(err)
}

func TestReadCSVOptionsPrettyPrint(t *testing.T) {
	assert := assert.New(t)

	file, err := os.Open("test/fixtures/data.csv")
	assert.NoError(err)

	res, err := ReadCSV(file, nil, Options{PrettyPrint: true})
	assert.NoError(err)

	json, err := ioutil.ReadFile("test/fixtures/pretty_print.json")
	assert.NoError(err)

	assert.Equal(strings.Trim(string(json), "\n"), string(res))
}
