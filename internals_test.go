package csv2json

import (
	"encoding/json"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

var (
	// TESTDATA used
	TESTDATA = [][]string{
		{"test", "data"},
		{"foo", "bar"},
		{"one", "two"},
	}
)

func TestIntoJSONArray(t *testing.T) {
	assert := assert.New(t)

	b, err := intoJSONArray(TESTDATA, nil, Options{})
	assert.NoError(err)

	var jsmap []map[string]string
	err = json.Unmarshal(b, &jsmap)
	assert.NoError(err)

	assert.Equal(2, len(jsmap))
	assert.Equal("foo", jsmap[0]["test"])
	assert.Equal("two", jsmap[1]["data"])
}

func TestIntoJSONArrayAdditionalFields(t *testing.T) {
	assert := assert.New(t)

	addFields := make(map[string]string)
	addFields["addedField"] = "test"

	b, err := intoJSONArray(TESTDATA, addFields, Options{})
	assert.NoError(err)

	var jsmap []map[string]string
	err = json.Unmarshal(b, &jsmap)
	assert.NoError(err)

	assert.Equal(2, len(jsmap))
	assert.Equal("test", jsmap[0]["addedField"])
	assert.Equal("test", jsmap[1]["addedField"])
}

func TestLineWiseJSON(t *testing.T) {
	assert := assert.New(t)

	b, err := lineWiseJSON(TESTDATA, nil, Options{})
	assert.NoError(err)

	for _, js := range strings.Split(string(b), "\n")[0:1] {
		var jsmap map[string]string
		err = json.Unmarshal([]byte(js), &jsmap)
		assert.NoError(err)
		assert.NotEmpty(jsmap)
	}
}

func TestLineWiseJSONAdditionalFields(t *testing.T) {
	assert := assert.New(t)

	addFields := make(map[string]string)
	addFields["addedField"] = "test"

	b, err := lineWiseJSON(TESTDATA, addFields, Options{})
	assert.NoError(err)
	for _, js := range strings.Split(string(b), "\n")[0:1] {
		var jsmap map[string]string
		err = json.Unmarshal([]byte(js), &jsmap)
		assert.NoError(err)
		assert.NotEmpty(jsmap)
		assert.Equal(jsmap["addedField"], "test")
	}
}
