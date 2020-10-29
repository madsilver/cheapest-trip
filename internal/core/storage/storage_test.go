package storage

import (
	"testing"

	"github.com/magiconair/properties/assert"
)

const msg = "got %v want %v given"

func TestLoadCsvNoFileGiven(t *testing.T) {
	args := []string{"main.go"}
	data, err := LoadCsv(args)

	if data != nil {
		t.Errorf(msg, data, nil)
	}
	assert.Equal(t, err.Error(), "No file given")
}

func TestLoadCsvNoSuchFile(t *testing.T) {
	args := []string{"main.go", "input-file.csv"}
	data, err := LoadCsv(args)

	if data != nil {
		t.Errorf(msg, data, nil)
	}
	assert.Equal(t, err.Error(), "open input-file.csv: no such file or directory")
}
