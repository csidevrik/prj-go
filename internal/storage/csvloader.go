// File: csvloader.go
// Package storage provides functionality to load and parse CSV files.
// It defines a Record structure to represent each row in the CSV file and a DataNet structure to hold a list of records.
// The LoadCSV function reads a CSV file from the specified filepath and returns a DataNet object containing the parsed records.
// It uses the csvutil package for decoding the CSV data into the Record structure.
// The Record structure contains fields for "name" and "IP", which are mapped to the corresponding CSV columns using struct tags.
package storage

import (
	"encoding/csv"
	"os"

	"github.com/jszwec/csvutil"
)

type Record struct {
	Name string `json:"name" csv:"name"`
	IP   string `json:"ip" csv:"IP"`
}

type DataNet struct {
	Records []Record `json:"datalist"`
}

func LoadCSV(filepath string) (DataNet, error) {
	f, err := os.Open(filepath)
	if err != nil {
		return DataNet{}, err
	}
	defer f.Close()

	r := csv.NewReader(f)
	var records []Record
	dec, _ := csvutil.NewDecoder(r)
	for {
		var rec Record
		err := dec.Decode(&rec)
		if err != nil {
			break
		}
		records = append(records, rec)
	}
	return DataNet{Records: records}, nil
}
