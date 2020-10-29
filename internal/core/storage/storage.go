package storage

import (
	"encoding/csv"
	"errors"
	"os"
)

// CheckCsv csv file
func CheckCsv(args []string) (f *os.File, err error) {
	if len(args) == 1 {
		err = errors.New("No file given")
		return
	}

	f, err = os.Open(args[1])
	if err != nil {
		return
	}

	return
}

// LoadCsv load data from csv file
func LoadCsv(args []string) (rows [][]string, err error) {
	f, err := CheckCsv(args)
	if err != nil {
		return
	}
	defer f.Close()

	reader := csv.NewReader(f)
	rows, err = reader.ReadAll()
	return
}

// WriteCsv writes data into csv file
func WriteCsv(args []string, data [][]string) (err error) {
	f, err := os.OpenFile(args[1], os.O_CREATE|os.O_WRONLY, 0777)
	if err != nil {
		return
	}
	defer f.Close()

	w := csv.NewWriter(f)
	defer w.Flush()

	w.WriteAll(data)

	return
}
